const navigationGroups = document.querySelectorAll('.nav-group');

navigationGroups.forEach((group) => {
  group.addEventListener('toggle', () => {
    if (!group.open) return;
    navigationGroups.forEach((other) => {
      if (other !== group) other.open = false;
    });
  });
});

document.addEventListener('click', (event) => {
  navigationGroups.forEach((group) => {
    if (!group.contains(event.target)) group.open = false;
  });
});

document.addEventListener('keydown', (event) => {
  if (event.key !== 'Escape') return;
  navigationGroups.forEach((group) => {
    if (!group.open) return;
    group.open = false;
    group.querySelector('summary').focus();
  });
});

const contactForm = document.querySelector('#contact-form');

contactForm?.addEventListener('submit', (event) => {
  event.preventDefault();
  const fields = new FormData(contactForm);
  const name = `${fields.get('firstName')} ${fields.get('lastName')}`;
  const body = `${fields.get('message')}\n\n${name}\n${fields.get('email')}`;
  const subject = encodeURIComponent(fields.get('subject'));
  window.location.href = `mailto:info@chopshopdance.org?subject=${subject}&body=${encodeURIComponent(body)}`;
});
