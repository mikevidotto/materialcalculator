window.JW_SHEDS_CONFIG = {
  // Your Go server already uses:
  // http.HandleFunc("/calculate", CalculateMaterials)
  //
  // Keeping this relative means it works whether the site is served from
  // localhost:8085, a domain name, or behind HTTPS on the same origin.
  calculateEndpoint: "/api/calculate"
};
