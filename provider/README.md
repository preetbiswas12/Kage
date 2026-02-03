# Providers

Builtin providers for manga sources.
They are faster and less memory consuming than the custom ones written in Lua.

## How Providers Connect to Manga Websites

### Built-in Providers

Mangal includes four built-in providers:

1. **Mangadex** - Uses official MangaDex API client
2. **Manganelo** - Web scraping with Colly
3. **Manganato** - Web scraping with Colly
4. **Mangapill** - Web scraping with Colly

### Connection Methods

#### API-Based (Mangadex)
- Uses the `github.com/darylhjd/mangodex` client library
- Communicates via official MangaDex REST API
- Provides structured data access with proper rate limiting

#### Web Scraping (Manganelo, Manganato, Mangapill)
- Uses `github.com/gocolly/colly/v2` for HTML parsing
- Extracts data using CSS selectors
- Three specialized collectors:
  - **Manga Collector**: Searches and extracts manga listings
  - **Chapter Collector**: Extracts chapter lists from manga pages
  - **Page Collector**: Extracts image URLs from chapter pages

### HTTP Client Configuration

All connections use the shared HTTP client from `network/client.go`:
- Request timeout: 1 minute
- Connection pooling: 100 idle connections, 100 per host
- Custom headers: User-Agent, Referer, Accept-Language
- Rate limiting: 50ms delay between requests (configurable)
- Parallelism: Up to 50 concurrent requests (configurable)

### For More Details

See [docs/SERVER_CONNECTION_ARCHITECTURE.md](../docs/SERVER_CONNECTION_ARCHITECTURE.md) for a comprehensive explanation of the connection architecture.