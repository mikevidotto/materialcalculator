# JW Sheds Frontend Recreation

This is a static, dependency-free recreation of the supplied JW Sheds screenshots.

## Pages

- `index.html`
- `prices.html`
- `features.html`
- `testimonials.html`
- `pictures.html`
- `contact.html`
- `privacy-policy.html`
- `quote.html`

The visual design is recreated from the screenshots supplied in the chat. Since this was built from screenshots rather than the original site's source files, some exact copy, imagery, prices, and spacing may differ from the live site.

## Shed dimension form

The Free Quote page includes a shed dimension form with:

- length in feet
- width in feet
- height in feet
- height defaults to 8 ft

It sends:

```http
POST /calculate
Content-Type: application/json
```

Example body:

```json
{
  "length": 12,
  "width": 8,
  "height": 8
}
```

The endpoint is configured in `config.js`:

```js
window.JW_SHEDS_CONFIG = {
  calculateEndpoint: "/calculate"
};
```

This is intentionally a relative URL. If your Go backend serves the frontend and the `/calculate` handler from the same host/port, this avoids CORS and avoids the production problem caused by using `localhost` in browser-side code.

## Recommended Go server

Assuming this whole folder is at `../../frontend` relative to the working directory of your Go process:

```go
package main

import (
    "log"
    "net/http"
    "path/filepath"
)

func main() {
    frontendPath, err := filepath.Abs("../../frontend")
    if err != nil {
        log.Fatal(err)
    }

    fs := http.FileServer(http.Dir(frontendPath))

    // API endpoint
    http.HandleFunc("/calculate", CalculateMaterials)

    // Website
    http.Handle("/", fs)

    log.Println("http://localhost:8085")
    log.Fatal(http.ListenAndServe(":8085", nil))
}
```

Then:

- `http://localhost:8085/` -> `index.html`
- `http://localhost:8085/prices.html`
- `http://localhost:8085/quote.html`
- `POST http://localhost:8085/calculate`

### If you insist on serving the frontend at `/static/`

Use:

```go
fs := http.FileServer(http.Dir(frontendPath))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

Then the site starts at:

`http://localhost:8085/static/`

The form still posts to `/calculate`, which goes to:

`http://localhost:8085/calculate`

## Before production

Replace or verify:

- pricing
- company contact information
- testimonial copy
- privacy policy wording
- map embed
- contact-form submission endpoint
- production analytics/cookie implementation
- reCAPTCHA if desired

The shed dimension form itself is already wired to `/calculate`.
