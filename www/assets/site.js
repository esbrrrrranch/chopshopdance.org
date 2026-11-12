const siteHeader = document.querySelector(".site-header");
const menuToggle = document.querySelector(".menu-toggle");
const mobileViewport = window.matchMedia("(max-width: 899px)");

const setMenuOpen = (open) => {
  siteHeader.dataset.menuOpen = String(open);
  menuToggle.setAttribute("aria-expanded", String(open));
  menuToggle.setAttribute(
    "aria-label",
    open ? "Close navigation" : "Open navigation",
  );
};

if (siteHeader && menuToggle) {
  setMenuOpen(false);
  menuToggle.addEventListener("click", () => {
    setMenuOpen(siteHeader.dataset.menuOpen !== "true");
  });
  mobileViewport.addEventListener("change", () => setMenuOpen(false));
  siteHeader.querySelector("nav").addEventListener("click", (event) => {
    if (event.target.closest("a")) setMenuOpen(false);
  });
}

const navigationGroups = document.querySelectorAll(".nav-group");

navigationGroups.forEach((group) => {
  group.addEventListener("toggle", () => {
    if (!group.open) return;
    navigationGroups.forEach((other) => {
      if (other !== group) other.open = false;
    });
  });
});

document.addEventListener("click", (event) => {
  if (siteHeader && !siteHeader.contains(event.target)) setMenuOpen(false);
  navigationGroups.forEach((group) => {
    if (!group.contains(event.target)) group.open = false;
  });
});

document.addEventListener("keydown", (event) => {
  if (event.key !== "Escape") return;
  if (mobileViewport.matches && siteHeader?.dataset.menuOpen === "true") {
    setMenuOpen(false);
    navigationGroups.forEach((group) => {
      group.open = false;
    });
    menuToggle.focus();
    return;
  }
  navigationGroups.forEach((group) => {
    if (!group.open) return;
    group.open = false;
    group.querySelector("summary").focus();
  });
});

const contactForm = document.querySelector("#contact-form");

contactForm?.addEventListener("submit", (event) => {
  event.preventDefault();
  const fields = new FormData(contactForm);
  const name = `${fields.get("firstName")} ${fields.get("lastName")}`;
  const body = `${fields.get("message")}\n\n${name}\n${fields.get("email")}`;
  const subject = encodeURIComponent(fields.get("subject"));
  window.location.href = `mailto:info@chopshopdance.org?subject=${subject}&body=${encodeURIComponent(body)}`;
});
