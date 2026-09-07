(() => {
  const page = document.body.dataset.page;

  document.querySelectorAll("[data-page-link]").forEach((link) => {
    if (link.dataset.pageLink === page) {
      link.classList.add("active");
      link.setAttribute("aria-current", "page");
    }
  });

  const toggle = document.querySelector("[data-nav-toggle]");
  const menu = document.querySelector("[data-nav-menu]");

  if (toggle && menu) {
    toggle.addEventListener("click", () => {
      const open = menu.classList.toggle("open");
      toggle.setAttribute("aria-expanded", String(open));
    });
  }

  document.querySelectorAll("[data-year]").forEach((node) => {
    node.textContent = new Date().getFullYear();
  });

  // Cookie notice — functional, but intentionally lightweight.
  const cookieCard = document.querySelector("[data-cookie-card]");
  const cookieChoice = localStorage.getItem("jw-cookie-choice");

  if (cookieCard && cookieChoice) {
    cookieCard.hidden = true;
  }

  document.querySelectorAll("[data-cookie-choice]").forEach((button) => {
    button.addEventListener("click", () => {
      localStorage.setItem("jw-cookie-choice", button.dataset.cookieChoice);
      if (cookieCard) cookieCard.hidden = true;
    });
  });

  // Shed dimension POST integration.
  const dimensionForm = document.querySelector("#shed-dimension-form");
  const dimensionStatus = document.querySelector("#dimension-status");

  if (dimensionForm) {
    const endpoint = window.JW_SHEDS_CONFIG?.calculateEndpoint || "/api/calculate";
    const endpointNode = document.querySelector("[data-calculate-endpoint]");
    if (endpointNode) endpointNode.textContent = endpoint;

    dimensionForm.addEventListener("submit", async (event) => {
      event.preventDefault();

      const submit = dimensionForm.querySelector('button[type="submit"]');
      const formData = new FormData(dimensionForm);

      const payload = {
        length: Number(formData.get("length")),
        width: Number(formData.get("width")),
        height: Number(formData.get("height"))
      };

      if (
        !Number.isFinite(payload.length) ||
        !Number.isFinite(payload.width) ||
        !Number.isFinite(payload.height) ||
        payload.length <= 0 ||
        payload.width <= 0 ||
        payload.height <= 0
      ) {
        showDimensionStatus("Please enter valid dimensions greater than zero.", "error");
        return;
      }

      submit.disabled = true;
      submit.textContent = "SENDING...";
      showDimensionStatus("Sending shed dimensions to the calculator...", "");

      try {
        const response = await fetch(endpoint, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Accept": "application/json, text/plain, */*"
          },
          body: JSON.stringify(payload)
        });

        const text = await response.text();

        if (!response.ok) {
          throw new Error(text || `HTTP ${response.status}`);
        }

        let display = text;
        try {
          const parsed = text ? JSON.parse(text) : {};
          display = JSON.stringify(parsed, null, 2);
        } catch (_) {
          // Plain text response is fine too.
        }

        const suffix = display && display !== "{}"
          ? `\n\nServer response:\n${display}`
          : "";

        showDimensionStatus(
          `Shed size submitted: ${payload.length} ft × ${payload.width} ft × ${payload.height} ft.${suffix}`,
          "success"
        );
      } catch (error) {
        showDimensionStatus(
          `Could not submit the dimensions to ${endpoint}.\n\n${error.message}\n\nMake sure your Go server has http.HandleFunc("/api/calculate", CalculateMaterials) and that this frontend is being served by the same server.`,
          "error"
        );
      } finally {
        submit.disabled = false;
        submit.textContent = "SUBMIT SHED SIZE";
      }
    });
  }

  function showDimensionStatus(message, state) {
    if (!dimensionStatus) return;
    dimensionStatus.hidden = false;
    dimensionStatus.className = "dimension-status" + (state ? ` ${state}` : "");
    dimensionStatus.textContent = message;
  }

  // Simple gallery interaction. Clicking a thumbnail moves it into the centre.
  const galleryCenter = document.querySelector("[data-gallery-center]");
  const galleryButtons = document.querySelectorAll("[data-gallery-thumb]");

  if (galleryCenter && galleryButtons.length) {
    galleryButtons.forEach((button) => {
      button.addEventListener("click", () => {
        const src = button.querySelector("img")?.getAttribute("src");
        if (!src) return;

        galleryCenter.setAttribute("src", src);
        galleryButtons.forEach((item) => item.classList.remove("active"));
        button.classList.add("active");
      });
    });
  }

  // Placeholder contact form behavior so it doesn't silently navigate away.
  document.querySelectorAll("[data-placeholder-form]").forEach((form) => {
    form.addEventListener("submit", (event) => {
      event.preventDefault();
      const note = form.querySelector("[data-form-note]");
      if (note) {
        note.hidden = false;
        note.textContent = "Contact form layout is ready. Connect this form to your preferred email or CRM endpoint when you are ready.";
      }
    });
  });
})();
