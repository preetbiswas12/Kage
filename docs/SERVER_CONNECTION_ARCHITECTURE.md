# Server Connection Architecture

## Overview

This document explains how Mangal (Kage) connects to manga websites to scrape and download manga content. The architecture is designed for performance, anti-detection, and extensibility.

## Network Layer

### HTTP Client Configuration
**Location:** `/network/client.go`

Mangal uses Go's standard `net/http` library with a custom-configured transport layer for optimal performance:

```go
var Client = &http.Client{
    Timeout:   time.Minute,
    Transport: transport,
}
```

### Connection Pooling

The HTTP transport is configured with aggressive connection pooling to minimize latency:

| Setting | Value | Purpose |
|---------|-------|---------|
| `MaxIdleConns` | 100 | Total idle connections across all hosts |
| `MaxIdleConnsPerHost` | 100 | Idle connections per domain |
| `MaxConnsPerHost` | 200 | Maximum concurrent connections per domain |
| `IdleConnTimeout` | 30 seconds | Closes idle connections after 30s |
| `ResponseHeaderTimeout` | 30 seconds | Timeout for reading response headers |
| `ExpectContinueTimeout` | 30 seconds | Timeout for 100-continue responses |

**Benefits:**
- Reuses TCP connections across multiple requests
- Reduces handshake overhead
- Improves download speed for manga pages
- Handles concurrent downloads efficiently

## Scraping Layer

### Colly Web Scraper
**Location:** `/provider/generic/new.go`

Mangal uses the [Colly](https://github.com/gocolly/colly) framework for web scraping. Colly provides:
- HTML parsing with goquery selectors
- Automatic retries and error handling
- Built-in caching
- Rate limiting
- Concurrent request management

### Three-Tier Collector System

Mangal uses three separate Colly collectors, each optimized for a specific task:

#### 1. Manga Search Collector
- **Purpose:** Search for manga titles and list results
- **Referer Header:** `https://google.com`
- **Rate Limiting:** Configurable per source

#### 2. Chapters Collector
- **Purpose:** Fetch chapter lists for a specific manga
- **Referer Header:** Manga URL (dynamic)
- **Rate Limiting:** Configurable per source

#### 3. Pages Collector
- **Purpose:** Extract page image URLs from a chapter
- **Referer Header:** Chapter URL (dynamic)
- **Rate Limiting:** Configurable per source

### Collector Configuration

```go
collectorOptions := []colly.CollectorOption{
    colly.AllowURLRevisit(),        // Re-visit URLs for caching
    colly.Async(true),               // Enable async mode
    colly.CacheDir(where.Cache()),   // Cache responses
}

baseCollector.SetRequestTimeout(20 * time.Second)
```

## Anti-Detection Measures

### Request Headers

All requests include browser-like headers to avoid detection as a bot:

```
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36
Accept: text/html
Accept-Language: en-US
Host: <source-domain>
Referer: <context-appropriate-url>
```

**Dynamic Referer Strategy:**
- Search requests → `https://google.com`
- Chapter requests → Manga URL
- Page requests → Chapter URL
- Image downloads → Chapter URL

This mimics natural browser navigation patterns.

### Rate Limiting

Configured per source via `colly.LimitRule`:

```go
colly.Limit(&colly.LimitRule{
    Parallelism: int(config.Parallelism),  // Concurrent requests
    RandomDelay: config.Delay,              // Delay between requests
    DomainGlob:  "*",                       // Apply to all domains
})
```

**Example (Manganelo):**
- Delay: 50ms between requests
- Parallelism: 50 concurrent goroutines
- Random jitter added automatically

### Caching

Responses are cached to disk at `where.Cache()` location:
- Reduces redundant requests
- Improves performance on repeated queries
- Respects source server resources

## Provider System

### Built-in Providers
**Location:** `/provider/init.go`

Four built-in sources are supported:

1. **Mangadex** - Uses external library with its own client
2. **Manganelo** - Generic scraper
3. **Manganato** - Generic scraper
4. **Mangapill** - Generic scraper

### Generic Provider Configuration
**Location:** `/provider/generic/configuration.go`

Each generic provider defines:
- Base URL
- CSS selectors for manga/chapter/page extraction
- Rate limiting parameters
- Header configurations

### Custom Lua Providers

Users can create custom sources using Lua scripts:
- Embedded Lua 5.1 VM (gopher-lua)
- Access to HTTP client, HTML parser, and headless Chrome
- Stored in `mangal where --sources` directory
- Hot-reloadable without recompiling

## Page Download Flow

**Location:** `/source/page.go`

The actual manga image download process:

1. **Create HTTP Request**
   ```go
   req, err := http.NewRequest(http.MethodGet, p.URL, nil)
   req.Header.Set("Referer", p.Chapter.URL)
   req.Header.Set("User-Agent", constant.UserAgent)
   ```

2. **Execute Request**
   ```go
   resp, err := network.Client.Do(req)  // Uses pooled client
   ```

3. **Validate Response**
   - Check HTTP 200 status
   - Verify content length > 0

4. **Buffer Content**
   - Read into memory buffer
   - Handle unknown content-length responses

5. **Process Image**
   - Detect image format
   - Store for export (PDF, CBZ, ZIP, or plain images)

## Architecture Summary

```
User Query
    ↓
┌───────────────────────┐
│   Provider System     │
│  (Built-in + Lua)     │
└───────────┬───────────┘
            ↓
┌───────────────────────┐
│   Colly Scrapers      │
│  (3-tier collectors)  │
└───────────┬───────────┘
            ↓
┌───────────────────────┐
│   HTTP Client         │
│  (Connection Pool)    │
└───────────┬───────────┘
            ↓
    Manga Website
```

## Key Design Decisions

✅ **Connection Pooling** - Reuse TCP connections for speed
✅ **Realistic Headers** - Mimic browser behavior
✅ **Dynamic Referers** - Follow natural navigation patterns
✅ **Rate Limiting** - Respect source servers
✅ **Caching** - Reduce redundant requests
✅ **Async Operations** - Parallel downloads
✅ **Extensibility** - Lua-based custom sources
✅ **No Proxy Support** - But configurable via environment variables

## Performance Characteristics

- **Concurrent Downloads:** Up to 200 connections per host
- **Request Timeout:** 20-60 seconds
- **Cache:** Disk-based, unlimited retention
- **Memory Usage:** Buffers images in memory during processing
- **CPU Usage:** Minimal (mostly I/O bound)

## Security Considerations

- User agent rotation: Single static user agent (could be enhanced)
- IP rotation: Not implemented (use proxy via environment)
- Cookie handling: Automatic via http.Client
- TLS verification: Enabled by default

## Future Enhancements

Potential improvements to the connection architecture:
- Multiple user agent rotation
- Proxy pool support
- Cloudflare bypass integration
- Custom DNS resolution
- Request fingerprint randomization
