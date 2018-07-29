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
      .match(event.request)
      .then(response => {
        if (response) {
          return response;
        }

        const fetchRequest = event.request.clone();

        return fetch(fetchRequest).then(response => {
          if (
            !response ||
            response.status !== 200 ||
            response.type !== 'basic'
          ) {
            return response;
          }

          const respondToCache = response.clone();

          caches.open(CACHE_NAME).then(cache => {
            cache.put(event.request, responseToCache);
          });

          return response;
        });
      })
      .catch(err => console.log('Error', err))
  );
});
