Post every day


## 11/1/2024 8:49:39 p. m.
![alt text](image.png)



Greetings everyone so I posted earlier that I was looking for an extension similar to pub manager that can give me an overiew of my project dependencies and groups apps according to projects. I could not find such a project, so I am building one called libTracker, everyone could suggest names if they like, not planning on doing a waitlist sign up becuase I have to have this tool to manage several projects. It will be freenium mabye free up to 3 projects and 2 apps per project, if that does not get conversion then I will increase the tier. Let me know your thoughts everyone


## 11/2/2024 7:23:35 p. m.


Greetings everyone working on my LibTracker app. Simplify dependency management with this simple SBOM Tool. so update I grouped the logic by projects and by apps so you can add projects and add apps. On the the project detail page you can view the apps in the project and the framework and programming language metadata for each app. Right now I only obtain project metadata for Angular.
What I learned
- I have to communicate from the panel js to the extension js if I want to have the input for the appPath be the path to the app. cant use webkitdirectory attribute for input type as it will take all files
- I learned package-lock.json contains the exact version of a package which is the package currently installed for the app and I should try to reference against that instead
Next goals
- CRUD operations for projects and apps
- obtain framework and programming language metadata for all apps listed in the popular section
- ability to open the terminal and or vscode window for each app
  - (mabye) mabye allow for the option to select a vscode workspace file if none exists
- (mabye) if the exact version of a dependencies cannot be determined for a framework or programming language then display a warning
- (backlog) generate an SBOM for each project and app





## 11/3/2024 1:26:30 p. m.
Greetings everyone working on my LibTracker app. Simplify dependency management with this simple SBOM Tool. So I did a lot CRUD for projects and apps. As well as opening the terminal and or vscode window for each app. I am not sure how to do the workspace logic for the project. The extension is now in the vscode extension marketplace you can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

So far I got Flutter android,ios,and macos support and I would need someone to explain where to find at least the platform version for windows and linux versions

# Learned
cant use alert,confirm, modals in webview panels
panel development is like one side is a the client and the code where the extension api lives is a server
coding in vanilla js is a great way to do tdd i suppose
there is npm libaries to parse gradle files!

# Next goals
- (main goal)obtain framework and programming language metadata for all apps listed in the popular section
- app dependency page
  - view licenses
  - view security vulnerabilities
  - crud on dependencies
  - (mabye) show packages not in use in the app
    - very hard would rely on executables

Let me know your thoughts





## 11/5/2024 3:54:14 p. m.
Greetings everyone working on my LibTracker app. Simplify dependency management with this simple SBOM Tool. Not too many updates just added support for selenium java and corrected a bug where it show 1.1.1 for all versions

The extension is now in the vscode extension marketplace you can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

# What I learned
npm has parsers for the dep files for many frameworks. Chatgpt can easily write code to extract crucial version information from the dep files as well as tell you where to look

# Next goals
- option to edit the framework type
- (main goal)obtain framework and programming language metadata for all apps listed in the popular section
- app dependency page
  - view licenses
  - view security vulnerabilities
  - crud on dependencies
  - (mabye) show packages not in use in the app
    - very hard would rely on executables

Let me know your thoughts



## 11/8/2024 2:59:52 p. m.
Greetings everyone working on my LibTracker Vscode Extension. Simplify dependency management with this simple SBOM Tool. Major updates supported all the popular frameworks along with their gradle, kotlin gradle, npm, pip and venv equivalents for java,node.js and python apps.

Since the package manager was important to the dependency info in the app I added it in as an additional when you go to add your app to your project.

I made a parser to parse Gemfile files, hopefully it can work with ios PodFiles and use it all to create a npm library to parse ruby related dep files

The extension is now in the vscode extension marketplace you can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

# Next goals
- app dependency page
  - view licenses
  - view security vulnerabilities
  - crud on dependencies
  - (mabye) show packages not in use in the app
    - very hard would rely on executables
- support for a code workspace so you can open the whole project as a code workspace
  - add support to edit that code-workspace files if there are additional non app folders that need to be in the workspace

Let me know your thoughts


## 11/11/2024
Greetings everyone working on my LibTracker Vscode Extension. Get to know your apps with this simple SBOM Tool. View at a a glance and fix outdated versions, security vulnerabilities and problematic licensing.
So not too many updates just support for ios apps and drag and drop functionality to reorder your apps

The extension is now in the vscode extension marketplace you can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.


## Next Goals
- app dependency page
  - view licenses
  - view security vulnerabilities
  - crud on dependencies
  - (mabye) show packages not in use in the app
    - very hard would rely on executables
- support for a code workspace so you can open the whole project as a code workspace
  - add support to edit that code-workspace files if there are additional non app folders that need to be in the workspace

Let me know your thoughts


## 11/13/24
* Greetings everyone working on my LibTracker Vscode Extension. Get to personally know your apps with this simple SBOM Tool.  View at a a glance and fix outdated versions, security vulnerabilities and problematic licensing.
* We are now over 100 users since our launch on Nov 5
* You can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.
* MAJOR UPDATE APP DETAIL PAGE
* we now get dependency info, version info and license info about each app in your project
* progress info from the app dependency page

# Next Goals
- focus on npm apps as the template for the rest of the apps
  - get security info
  - add a package
  - choose a version choose latest version
  - snapshot in case decision broke the app/ failed test
  - update all packages
  - remove packages
- bulk import bulk append bulk export
- refresh logic
- cache detailed dep http api calls
- coming soon notifications for buttons with no functionality


Let me know your thouhgts and tools suggest for security sources to get security vulnerabiltiy info about packages? mabye snyk
who will sponsor me snyk SBOM tools


# 11/16/24
*  Greetings everyone working on my LibTracker Vscode Extension. Get to personally know your apps with this simple SBOM Tool.  View at a a glance and fix outdated versions, security vulnerabilities and problematic licensing.
* You can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

## New Features
* import/export functionalitty
  * append, overwrite ,merge (if project name are the same in the file and extension and those apps are the same merge the two)
  * import validation
* support for all (not just popular) npm apps

## Next goals
- app detail page
  * security page
    - add a package
  - choose a version choose latest version
  - snapshot in case decision broke the app/ failed test
  - update all packages
  - remove packages
- refresh logic
- cache detailed dep http api calls

# 11/18/24
*  Greetings everyone working on my LibTracker Vscode Extension. Get to personally know your apps with this simple SBOM Tool.  View at a a glance and fix outdated versions, security vulnerabilities and problematic licensing.
* You can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

## New Features
* added support for the Analog.js Framework
* click on the latest version will take you to the change log for that package
* clicking on the package name will take you to the package homepage
* there are now dependency subtables and the table controls affect each subtable seperately instead of concerning the whole table in the table control action
* can choose which version you want to update to and view when each version was released for better cross refercing
* CRUD with packages
  * update/bulk update
  * add
  * delete
* refresh logic for all dashboard pages
* can backup and restore your package manager files and decide to automatically or manually resolve your packages
* offline cache support for http calls
  * enable,disable,toggle the cache

## Next goals
- sort functionality (Done 21 min)
- hide projects (this is for doing demonstrations) (Done 33 mins)
- open a terminal and run package management commands in the terminal so viewers know its working (55 mins) **issue** terminal.shellIntegration comes up undefined but falls back to using child process
  - add package
  - update all packages
  - install from backup
  - update
  - delete
- root folder for project so projects can be used across teams (18:55)
- sub dependencies of which the user doesnt know about (lock package amanger files aka the dependency tree)
- prefetch all of the users extension data
- app detail page
  - cve info
- Generate SBOM
- AI recommendations on what versions to use
- responsiveness app detail page can tabulator turn to a series of cards


## Lofy Milestones I might not do
- use git for backups to improve backup size
- compability table to see which combination of packages best works with your workflow. meaning users reporting that this works with that and so on
  - a notes section would be better I suppose
  - a wikipedia of people contributing to what works with what
- instead of refreshing the whole page refresh only the parts that changed
- see which packages are not in use in your app


# 11/23/24
*  Greetings everyone working on my LibTracker Vscode Extension. Get to personally know your apps with this simple SBOM Tool.  View at a a glance and fix outdated versions, security vulnerabilities and problematic licensing.
* You can access here https://marketplace.visualstudio.com/items?itemName=windmillcode-publisher-0.lib-tracker.

# New Features
- sort functionality
- hide projects (this is for doing demonstrations) (this does not work times as its more of an issue with vscode API)
- group paths to where you can assoicate multiple apps with a root folder to make it easier to transfer projects between computers
- app detail page
  - sub dependency info
    - if a dependency does not have a version it means it was never installed but it was optional unsure whether to include version range
- can view license information about subdependencies now
- toggle all sub deps and root deps for license and cve info
- when doing bulk group path update
  - ensure that it recursively looks for each app basename in the directory
  - ensure it knows the differnce between apps with an individual path and apps with a group path to ignore them
- app detail page
  - latest info , release date , package and changelog urls for sub deps
- prefetch all of the users extension data


# Next Goals
- recursion exclusion list
- (done) toggle select all apps in project detail page
- (mabye) workspace folder
  - (depends on capabilbility of vscode api to access vscode profiles)
- git backup changes
- app detail page
  - cve info
  - search (root row is possible but useless search every nested child row)
  - expanded and collapse all subtables in a given row
  - responsiveness app detail page can tabulator turn to  series of cards
- Generate SBOM
- (Question should we get license and cve info about every version of  every package of the app)
- URL or AI summary of categories and names for licenses and CVES
- (if possible) click on subdependency in license pane will take you to its location in table

