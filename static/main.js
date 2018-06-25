//-----------------------------------------
//
// Main
//
//-----------------------------------------

const LAZIES = document.getElementsByClassName('lazy');

let WIN_HEIGHT = window.innerHeight;

const deferWork = fn => {
  if (typeof requestIdleCallback !== 'undefined') {
    window.requestIdleCallback(fn, { timeout: 60 });
  } else if (typeof requestAnimationFrame !== 'undefined') {
    window.requestAnimationFrame(fn);
  } else {
    window.setTimeout(fn, 16.66);
  }

  return true;
};

function load(t) {
  const img = new Image();

  img.onload = () => {
    t.src = img.src;
    t.classList.add('loaded');
  };
  img.onerror = err => console.warn('err', err);

  img.src = t.getAttribute('data-src');
}

function onScroll() {
  deferWork(() => {
    for (let i = 0; i < LAZIES.length; i++) {
      const t = LAZIES[i];
      const { top } = t.getBoundingClientRect();
      if (top <= WIN_HEIGHT && !t.classList.contains('loaded')) {
        load(t);
      }
    }
  });
}

function onResize() {
  deferWork(() => {
    WIN_HEIGHT = window.innerHeight;
  });
}

function start() {
  const now = new Date();
  if (now.getDay() === 5) {
    document.body.classList.add('friday');
  }
  window.addEventListener('scroll', onScroll, { passive: true });
  window.addEventListener('resize', onResize, { passive: true });
}

if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/static/sw.js').then(
      registration => {},
      err => {
        console.log('Service worker failed to register', err);
      }
    );
  });
}

document.addEventListener('DOMContentLoaded', start);
