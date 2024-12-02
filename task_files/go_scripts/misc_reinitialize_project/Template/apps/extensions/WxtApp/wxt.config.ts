import { defineConfig } from 'wxt';


let config = {

  modules: ['@wxt-dev/i18n/module'],
  manifest:{
    default_locale: 'en',
    name:"AI Prompt Modifier",
    permissions: ["activeTab", "storage", "scripting"],
    host_permissions: ['*://*.chatgpt.com/*'],
    oauth2: {
      client_id: "modify-chatgpt-prompts-prod.apps.googleusercontent.com",
      scopes: [
        "email",
        "profile"
      ]
    },
    browser_specific_settings:{
      gecko: {
        id: "[GET_THIS_WHEN_CREATING_A_NEW_ADDON_IN_THE_MOZILLA_ADDON_STORE]@[PROJECT_NAME].com"
      }
    },
    content_security_policy: {
      "extension_pages": "script-src 'self'; object-src 'self'; img-src 'self' blob: data: unsafe-eval"
    },

    description: "[WEB_SEO_DESCRIPTION]"
  },
  runner:{
    keepProfileChanges: true,
    chromiumArgs: [
      "--new-window=https://chatgpt.com"
    ]
  }
}
// See https://wxt.dev/api/config.html
export default defineConfig(config);
