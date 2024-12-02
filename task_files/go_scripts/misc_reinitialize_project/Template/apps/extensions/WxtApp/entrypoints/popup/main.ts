import './style.css';
import { initializeApp } from "firebase/app";
import { connectAuthEmulator, getAuth, GoogleAuthProvider, createUserWithEmailAndPassword, onAuthStateChanged, signInWithEmailAndPassword, sendPasswordResetEmail } from "firebase/auth/web-extension";

const ENV_TYPE = import.meta.env.COMMAND ==="serve" ? "DEV" :"PROD"
// const ENV_TYPE = "DEV"
const HOST_COMPUTER:"LOCAL" |"REMOTE" = "LOCAL"
const IS_SAFARI = import.meta.env.SAFARI
const IS_FIREFOX = import.meta.env.FIREFOX
const IS_FREE = true
// SENTRY
import {BrowserClient,defaultStackParser,getDefaultIntegrations,makeFetchTransport,Scope} from "@sentry/browser";

import { getStorageItem, setStorageItem } from '@/public/utils';
// filter integrations that use the global variable
const integrations = getDefaultIntegrations({}).filter(
  (defaultIntegration) => {
    return !["BrowserApiErrors", "Breadcrumbs", "GlobalHandlers"].includes(
      defaultIntegration.name,
    );
  },
);


const client = new BrowserClient({
  environment:"Browser_Extension_"+ENV_TYPE.toUpperCase(),
  dsn: "[SENTRY_DSN]",
  transport: makeFetchTransport,
  stackParser: defaultStackParser,
  integrations: integrations,
  tracesSampleRate: ["DEV","TEST","DOCKER_DEV"].includes(ENV_TYPE)? 0.0: 1.0
});

const scope = new Scope();
scope.setClient(client);

client.init(); // initializing has to be done after setting the client on the scope
// SENTRY




let appTitle = document.querySelector<HTMLTitleElement>('title')!
appTitle.innerHTML = i18n.t("global.title");
let myApp = document.querySelector<HTMLDivElement>('#app')!


myApp.innerHTML = `
  <div>YOUR DATA HERE</div>
`;

class WMLUri {

  url: URL;
  // @ts-ignore
  constructor(props:any ={ scheme, host, port : null, path : '', query : '', fragment : '' }) {
    let { scheme, host, port, path, query, fragment } = props;
    this.url = new URL(`${scheme}://${host}${port ? `:${port}` : ''}`);
    if (path) this.url.pathname = path;
    if (query) this.url.search = query;
    if (fragment) this.url.hash = fragment;
  }

  get domain() {
    return this.url.port === '' ? this.url.hostname : `${this.url.hostname}:${this.url.port}`;
  }

  get fqdn() {
    return `${this.url.protocol}//${this.url.hostname}${this.url.port ? `:${this.url.port}` : ''}`;
  }

  toString() {
    return this.url.toString();
  }
}

class BackendUri extends WMLUri {

  constructor(props:any = { path: '', query: '', fragment: '' }) {
    let { path, query, fragment } = props;
    let host = {
      DEV: "example.com",
      PROD:"api.[PROJECT_NAME].com"
    }[ENV_TYPE]
    let port = {
      DEV: 10072,
      PROD: 443
    }[ENV_TYPE]
    if(ENV_TYPE === "DEV" && HOST_COMPUTER === "REMOTE"){
      host = "[PROXY_URLS_0]"
      port = 443
    }
    super({ scheme: 'https', host, port, path, query, fragment });
  }
}

const currentUser:any = {
  user: null,
  get accessToken() {
    return currentUser.user.accessToken
  },
  // should always have a length of 0 or 1 if there is more there is an error with your code
  subscriptions: []
}

const ENV :any= {
  type: ENV_TYPE,
  backendURL: new BackendUri(),
  app:{
    hostComputer:HOST_COMPUTER
  },
  firebase: {
    config: {
      apiKey: "",
      authDomain:"127.0.0.1",
      projectId: "[PROJECT_NAME]",
      storageBucket: "[PROJECT_NAME].appspot.com",
      messagingSenderId: "",
      appId: "",
      measurementId: ""
    },
    googleProvider: new GoogleAuthProvider()
  },
  // EXAMPLE WAYS TO MANAGE YOUR ENDPOINT YOU CAND DECIDE LATER
  accountsService: {
    listUsers: new BackendUri({
      path: 'accounts/users/list'
    }),
    deleteUser: new BackendUri({
      path: 'accounts/users/delete'
    })
  },
  storeService: {
    purchaseSubuscription: new BackendUri({
      path: '/store/subscriptions/purchase'
    }),
    cancelSubscription: new BackendUri({
      path: '/store/subscriptions/cancel'
    })
  }
}

initializeApp(ENV.firebase.config)
ENV.firebase.auth = getAuth(ENV.firebase.app)
ENV.firebase.auth.useDeviceLanguage()
if (ENV.type === "DEV") {
  if(ENV.app.hostComputer === "REMOTE"){
    ENV.firebase.config.authDomain = "[THE IP TO THE HOST COMPUTER]"
  }
  connectAuthEmulator(
    ENV.firebase.auth,
    `http://${ENV.firebase.config.authDomain}:[Firebase_Emulator_Auth_0]`,
    { disableWarnings: true }
  )

}
if (ENV.type === "PROD") {
  ENV.firebase.config.authDomain = "[PROJECT_NAME].firebaseapp.com"
}


function safeParseInt(value:string, radix = 10, myDefault = 1) {
  const result = parseInt(value, radix);
  return isNaN(result) ? myDefault : result;
}

function debounce(func:Function, wait:number) {
  let timeout:any;
  return function (...args:any) {
    clearTimeout(timeout);
    // @ts-ignore
    timeout = setTimeout(() => func.apply(this, args), wait);
  };
}


document.addEventListener('DOMContentLoaded', async () => {

  let tabs = await browser.tabs.query({ active: true, currentWindow: true })
  let activeTabId = tabs[0].id as number

});


