# Server Connection Architecture

## Overview

Mangal connects to manga websites through a multi-layered architecture that supports both built-in providers and custom Lua scrapers. This document explains how the application establishes connections and retrieves manga data from various sources.

## Connection Architecture

### 1. HTTP Client Layer (`network/client.go`)

At the foundation is a configured HTTP client that handles all network requests:

```go
var Client = &http.Client{
    Timeout:   time.Minute,
    Transport: transport,
}
```

**Configuration:**
- **Connection pooling**: Max 100 idle connections, 100 per host, 200 max per host
- **Timeouts**: 
  - Overall request timeout: 1 minute
  - Idle connection timeout: 30 seconds
  - Response header timeout: 30 seconds
  - Expect continue timeout: 30 seconds

This client is used for downloading manga page images and is the base transport mechanism.

### 2. Provider System

Mangal supports two types of providers:

#### A. Built-in Providers

**Four built-in manga sources:**
1. **Mangadex** - Uses the official MangaDex API client
2. **Manganelo** - Web scraping using Colly
3. **Manganato** - Web scraping using Colly  
4. **Mangapill** - Web scraping using Colly

##### Mangadex Connection (API-based)

Located in `provider/mangadex/mangadex.go`:

```go
type Mangadex struct {
    client *mangodex.DexClient
    cache  struct {
        mangas   *cacher[[]*source.Manga]
        chapters *cacher[[]*source.Chapter]
    }
}
```

- Uses the `github.com/darylhjd/mangodex` client library
- Communicates via official MangaDex REST API
- Includes caching for manga and chapter data

##### Generic Web Scraping (Colly-based)

For Manganelo, Manganato, and Mangapill (`provider/generic/new.go`):

**Technology:** Uses `github.com/gocolly/colly/v2` for web scraping

**Features:**
- Asynchronous scraping with configurable parallelism
- Disk-based caching in the user's cache directory
- Rate limiting with configurable delays (50ms default)
- Custom request headers:
  ```go
  r.Headers.Set("Referer", "https://google.com")
  r.Headers.Set("accept-language", "en-US")
  r.Headers.Set("Accept", "text/html")
  r.Headers.Set("Host", s.config.BaseURL)
  r.Headers.Set("User-Agent", constant.UserAgent)
  ```

**Three specialized collectors:**
1. **Manga Collector** - Searches for manga using CSS selectors
2. **Chapter Collector** - Extracts chapter lists from manga pages
3. **Page Collector** - Extracts image URLs from chapter pages

**Configuration example** (`provider/manganelo/manganelo.go`):
```go
Config = &generic.Configuration{
    Name:            "Manganelo",
    Delay:           50 * time.Millisecond,
    Parallelism:     50,
    ReverseChapters: true,
    BaseURL:         "https://ww5.manganelo.tv/",
    GenerateSearchURL: func(query string) string { ... },
    MangaExtractor:    &generic.Extractor{ ... },
    ChapterExtractor:  &generic.Extractor{ ... },
    PageExtractor:     &generic.Extractor{ ... },
}
```

Each extractor uses CSS selectors and extraction functions to parse HTML and extract:
- Manga names, URLs, and cover images
- Chapter names, URLs, and volumes
- Page image URLs

#### B. Custom Providers (Lua Scrapers)

Located in `provider/custom/`:

**Features:**
- Users can write custom scrapers in Lua 5.1
- Built-in Lua VM using `github.com/yuin/gopher-lua`
- Access to specialized libraries via `github.com/metafates/mangal-lua-libs`
- Supports headless Chrome for JavaScript-heavy sites

**Loading process** (`provider/custom/loader.go`):
1. Compile Lua source file to bytecode
2. Create new Lua state with preloaded libraries
3. Execute the script
4. Validate required functions exist
5. Return a source interface implementation

**Available capabilities in Lua:**
- HTTP client
- HTML parser
- Headless Chrome (when needed)
- JSON parsing
- And more...

### 3. Data Flow

#### Search Flow:
1. User enters search query
2. Provider generates search URL
3. HTTP request sent with proper headers
4. HTML response parsed using CSS selectors (or API response parsed)
5. Manga list extracted and returned

#### Chapter List Flow:
1. User selects a manga
2. Provider visits manga URL
3. Chapter list extracted from HTML
4. Chapters sorted (optionally reversed)
5. Chapter list returned with metadata

#### Page Download Flow:
1. User selects chapters to download
2. Provider visits chapter URL
3. Page URLs extracted from HTML
4. For each page (`source/page.go`):
   ```go
   func (p *Page) Download() error {
       req, err := http.NewRequest(http.MethodGet, p.URL, nil)
       req.Header.Set("Referer", p.Chapter.URL)
       req.Header.Set("User-Agent", constant.UserAgent)
       
       resp, err := network.Client.Do(req)
       // Download image data
   }
   ```
5. Images downloaded asynchronously or synchronously (configurable)
6. Images cached in memory
7. Converted to desired format (PDF, CBZ, ZIP, or plain images)

### 4. Connection Security & Best Practices

**Headers Set:**
- **User-Agent**: Custom user agent to identify the client
- **Referer**: Set to maintain proper referrer chain (Google for searches, manga URL for chapters, chapter URL for pages)
- **Accept-Language**: "en-US" for English content
- **Accept**: "text/html" for HTML content

**Rate Limiting:**
- Configurable delay between requests (default 50ms)
- Parallelism limits (default 50 concurrent requests)
- Domain-level rate limiting with Colly

**Caching:**
- Disk-based HTTP cache via Colly
- In-memory caching for manga/chapter data
- Reduces redundant network requests

**Error Handling:**
- HTTP status code validation
- Response content length validation
- Proper error propagation through the stack
- Retry logic at application level

## Summary

The connection architecture is designed for:
- **Flexibility**: Support both API-based and scraping-based sources
- **Performance**: Async operations, connection pooling, caching
- **Extensibility**: Custom Lua scrapers for any manga website
- **Reliability**: Proper headers, rate limiting, error handling
- **Maintainability**: Clean separation of concerns between network layer, provider layer, and business logic

The system uses standard HTTP/HTTPS connections with proper headers and respects rate limits to avoid overwhelming source websites.
