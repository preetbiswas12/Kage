# Server Connection Architecture

This document explains how Mangal connects to manga websites to fetch and download manga content.

## Overview

Mangal uses a multi-layered architecture to connect to manga websites:

1. **HTTP Client Layer** - Low-level HTTP connections
2. **Provider Layer** - Abstract interface for different manga sources
3. **Source Layer** - Concrete implementations for specific websites
4. **Download Layer** - Coordinates the download process

## HTTP Client Layer

### Network Client (`network/client.go`)

The base HTTP client is configured in `network/client.go`:

```go
var Client = &http.Client{
    Timeout:   time.Minute,
    Transport: transport,
}
```

**Connection Settings:**
- **Max Idle Connections**: 100
- **Max Idle Connections Per Host**: 100
- **Max Connections Per Host**: 200
- **Idle Connection Timeout**: 30 seconds
- **Response Header Timeout**: 30 seconds
- **Expect Continue Timeout**: 30 seconds
- **Request Timeout**: 1 minute

These settings are optimized for:
- High concurrency when downloading multiple pages
- Connection reuse to reduce overhead
- Reasonable timeouts to handle slow connections

## Provider System

Mangal supports two types of providers:

### 1. Built-in Providers

Located in `provider/` directory, these are native Go implementations:

- **Mangadex** - Uses the official MangaDex API client
- **Manganelo** - Uses the generic scraper
- **Manganato** - Uses the generic scraper  
- **Mangapill** - Uses the generic scraper

#### Generic Scraper (`provider/generic/`)

The generic scraper uses the Colly framework for web scraping:

**Key Features:**
- **HTML Parsing**: Uses goquery to parse HTML and extract data
- **Concurrent Requests**: Configured parallelism for faster scraping
- **Caching**: Caches responses to disk to reduce repeated requests
- **Rate Limiting**: Configurable delays between requests
- **Custom Headers**: Sets appropriate headers for each request

**Request Headers Set:**
```go
- Referer: (Previous page or google.com)
- Accept-Language: en-US
- Accept: text/html
- Host: (Target website domain)
- User-Agent: (Mangal user agent)
```

### 2. Custom Lua Providers

Located in user's sources directory (accessible via `mangal where --sources`):

**How They Work:**
1. Lua scripts are compiled using gopher-lua
2. Required functions must be defined: search, manga details, chapters, pages
3. Uses `mangal-lua-libs` for HTTP requests and HTML parsing
4. Can optionally use headless browser via rod

**Advantages:**
- Users can add support for any manga website
- No need to recompile Mangal
- Community can share scrapers
- Can handle JavaScript-heavy websites (with headless mode)

## Connection Flow

### 1. Search Request

```
User Query → Provider.Search() → HTTP GET to search URL → Parse HTML → Return Manga List
```

**For Generic Providers:**
- Generates search URL from query
- Makes HTTP GET request with configured headers
- Parses HTML using CSS selectors from configuration
- Extracts manga titles, URLs, and cover images

### 2. Fetching Chapters

```
Manga Selected → Source.ChaptersOf() → HTTP GET to manga page → Parse HTML → Return Chapter List
```

**Request Details:**
- Referer set to the manga's URL
- Parses chapter list using configured selectors
- Extracts chapter names, URLs, and volumes
- Results are cached to avoid repeated requests

### 3. Fetching Pages

```
Chapter Selected → Source.PagesOf() → HTTP GET to chapter page → Parse HTML → Return Page URLs
```

**Request Details:**
- Referer set to the chapter's URL
- Extracts image URLs for all pages
- Stores page index and extension information

### 4. Downloading Pages

```
Page List → Page.Download() → HTTP GET for each image → Save to buffer → Convert & Save
```

**Download Process:**
- Creates HTTP request with proper referer and user agent
- Uses the shared HTTP client from network package
- Downloads image data into memory buffer
- Tracks download size and progress
- Images are then converted to selected format (PDF, CBZ, ZIP, etc.)

## Request Optimization

### Connection Pooling
- Reuses TCP connections via `MaxIdleConns` settings
- Reduces handshake overhead for multiple requests
- Up to 200 concurrent connections per host

### Caching
- Generic scraper caches responses to disk
- Cache directory: `mangal where --cache`
- Reduces load on manga websites
- Speeds up repeated searches and browsing

### Rate Limiting
- Configurable delays between requests
- Domain-specific rate limiting via colly
- Prevents overwhelming manga websites
- Reduces chance of being blocked

### Parallel Downloads
- Multiple pages can be downloaded concurrently
- Configured via parallelism settings
- Balances speed with server courtesy

## MangaDex Special Case

MangaDex uses an official API client (`github.com/darylhjd/mangodex`) instead of web scraping:

- Communicates with MangaDex API v5
- Uses proper authentication if configured
- Respects API rate limits
- More reliable than scraping
- Gets structured JSON data instead of parsing HTML

## Security Considerations

### Headers
- User-Agent is set to identify Mangal
- Referer header mimics browser behavior
- Accept headers indicate HTML content expected

### Timeouts
- Request timeout prevents hanging connections
- Response header timeout prevents slow headers
- Idle timeout cleans up unused connections

### Error Handling
- HTTP errors are properly caught and logged
- Empty responses are detected and reported
- Network errors trigger retry logic where appropriate

## Troubleshooting

### Connection Issues

**Symptom**: Timeout errors
- **Cause**: Website is slow or down
- **Solution**: Increase timeout in config or try different source

**Symptom**: 403 Forbidden
- **Cause**: Website blocking automated requests
- **Solution**: Use different source or custom scraper with headless mode

**Symptom**: Empty results
- **Cause**: Website changed their HTML structure
- **Solution**: Update scraper configuration or use different source

### Custom Scrapers

Custom Lua scrapers can access:
- HTTP client library for requests
- HTML parsing library (goquery-like)
- Headless browser (rod) for JavaScript sites
- JSON parsing utilities
- String manipulation functions

See [mangal-scrapers repository](https://github.com/metafates/mangal-scrapers) for examples.

## Performance Tips

1. **Use Built-in Providers**: Faster and more efficient than Lua scrapers
2. **Enable Caching**: Reduces repeated requests significantly
3. **Adjust Parallelism**: Balance between speed and server load
4. **Monitor Network**: Use `--log-level trace` to see all requests
5. **Respect Rate Limits**: Avoid getting blocked by being too aggressive

## Related Files

- `network/client.go` - HTTP client configuration
- `provider/provider.go` - Provider system
- `provider/generic/` - Generic web scraper
- `provider/mangadex/` - MangaDex implementation
- `provider/custom/` - Lua scraper loader
- `source/page.go` - Page download logic
- `downloader/download.go` - Download orchestration
