//-----------------------------------------
//
// Service worker
//
//-----------------------------------------

const VERSION = '0.0.1';
const CACHE_NAME = `james-pants-${VERSION}`;

const URLS_TO_CACHE = ['/', '/index.html'];

self.addEventListener('install', event => {
  const timestamp = Date.now();
  event.waitUntil(
    caches
      .open(CACHE_NAME)
      .then(cache => {
        console.log('opened cache');
        return cache.addAll(URLS_TO_CACHE);
      })
      .then(() => self.skipWaiting())
  );
});

self.addEventListener('activate', event => {
  event.waitUntil(self.clients.claim());
});

self.addEventListener('fetch', event => {
  event.respondWith(
    caches
      .open(cacheName)
      .then(cache => cache.match(event.request, { ignoreSearch: true }))
      .then(response => {
        return response || fetch(event.request);
      })
  );
});
