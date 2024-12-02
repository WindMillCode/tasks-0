import { Routes } from '@angular/router';
import { addPageTitlesToRoute, removeLeadingSlash } from '@core/utility/route-utils';
import { ENV } from '@env/environment';



export let routes: Routes = [

  {
    path: ENV.nav.urls.siteOffline,
    loadComponent: () =>
      import('./core/site-offline/site-offline.component').then(
        (m) => m.SiteOfflineComponent,
      ),
  },
  {
    path: ENV.nav.urls.notFound,
    loadComponent: () =>
      import('./core/not-found/not-found.component').then(
        (m) => m.NotFoundComponent,
      ),
  },
  {
    path: '**',
    redirectTo:ENV.nav.urls.notFound
  },
];


if (['DEV', 'DOCKER_DEV'].includes(ENV.type)) {
  let scratchpadRoute = {
    path: "scratchpad",
    loadComponent: () =>
      import('./pages/scratchpad/scratchpad.component').then(
        (m) => m.ScratchpadComponent,
      ),
  };
  routes.splice(1, 0, scratchpadRoute);
}


routes = removeLeadingSlash(routes)
routes= addPageTitlesToRoute(routes)

