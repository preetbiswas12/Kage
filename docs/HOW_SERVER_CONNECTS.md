# How the Server Connects to Manga Websites

## Quick Answer

Mangal connects to manga websites using:

1. **HTTP Client with Connection Pooling** (`/network/client.go`)
   - Go's `net/http` standard library
   - 200 max concurrent connections per domain
   - 100 idle connections kept alive for reuse
   - 1 minute request timeout

2. **Colly Web Scraping Framework** (`/provider/generic/`)
   - Three-tier collector system (manga search, chapters, pages)
   - Browser-like HTTP headers to avoid detection
   - Rate limiting (50ms delay, 50 concurrent requests)
   - Disk-based response caching

3. **Dynamic Request Headers**
   - `User-Agent`: Mozilla/5.0 (Chrome on Windows)
   - `Referer`: Context-appropriate (Google for search, chapter URL for images)
   - `Accept`: text/html
   - `Accept-Language`: en-US

4. **Provider System**
   - 4 built-in sources (Mangadex, Manganelo, Manganato, Mangapill)
   - Extensible via Lua scripts (custom sources)

## Connection Flow

```
User Request
    ↓
[Provider System]
    ↓
[Colly Scraper] → Set browser headers, rate limit, cache
    ↓
[HTTP Client] → Connection pooling, timeout handling
    ↓
Manga Website (thinks it's a browser)
```

## Key Features

- ✅ Connection reuse for performance
- ✅ Browser-like headers to avoid bot detection
- ✅ Rate limiting to respect servers
- ✅ Automatic caching of responses
- ✅ Parallel downloads (async mode)

For detailed technical information, see [SERVER_CONNECTION_ARCHITECTURE.md](SERVER_CONNECTION_ARCHITECTURE.md).
