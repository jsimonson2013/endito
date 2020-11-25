# Endito
A service for editing static web pages requiring no database connection.

## Support
If you enjoy the service and want to support me in further open source development.

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=8P7E4L2D2K6BU)

# Content editing
CMS, or Content Management Systems are very common in today's web development environment. They allow website content updates to be frequent and seemless offering varying levels of freedom in edits. Many CMS have steep learning curves, specific tech stacks, and complex configurations. At their core, any content editing system must have a few components
- awareness of the source files that need to be updated
- input system to receive file update requests
- a translation of input to file schema or syntax

## Why Endito?
If you want a simple content management system that supports seemless user editing, Endito may be the right system for you. It has few dependencies, minimal configuration, and is extremely portable.

# Usage
Endito is run as a companion service on the same host as the HTML file server. It accepts a path to the HTML file root directory on startup and can fetch auth from a variety of ways. It receives a POST request from the user interface, performs a git commit, and updates the file.

## How it works
An HTML file can be fetched and Endito renders it with a WYSIWYG-style wrapper that leverages [contenteditable attribute](https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/Editable_content). After edits are performed through Endito's editor, the user can enter their credentials to submit the edit request. If the HTML file's host server supports Endito, and the user's authorization is successful, the file will be uploaded, changes will be commited/pushed, and the new HTML files will be present.

## Setup
The benefit of Endito is that setup is simple. It's outlined in the steps below.

1. clone the endito repository
2. set the environment variables listed below
3. run the tool with `go run main.go` from the `endito/tool` folder
4. access the editor from `<host>:<port>/path-to/endito/editor`
5. load pages, edit things, post updates!

## HTML prep
In order for Endito's editor to be aware of HTML content that is editable, an class attribute must be added to the HTML element. Add these classes to the HTML elements you want to edit in your HTML files. For example,
```html
<h1 class="edit-text">Welcome to My Cool Website!</h1>
```
Here are the current classes that Endito's editor supports:
| Class | Description | Detail |
| ----- | ----------- | ------ |
| edit-text | used to signal text that can be edited | Applies highlights to element in editor, supports textual updates including copy+paste |

Additionally, to ensure that your styles are applied within the Endito editor, use absolute paths to your CSS files. For example,
```html
<link rel="stylesheet" href="file://///wsl$/Ubuntu-18.04/www/var/html/endito/example/home.css">
```

*NOTE: if there is a class you think endito needs, feel free to make a PR with that class*

## Environment Variables
Endito makes use of environment variables to access secret information and environment-specific information. For example, Endito uses `$BASE_DIR` and `$RLTV_DIR` to determine location of source html files and git project.

| Var | Definition | Example |
| --- | ---------- | ----------- |
| USERNAME | credential for endito user auth | test |
| PASSWORD | credential for endito user auth | pass |
| GIT_UNAME | git config for commit messages | gituser |
| GIT_EMAIL | git email for commit messages | user@github.com |
| BASE_DIR | absolute path to project root | /www/var/html |
| RLTV_DIR | relative path to project root from endito tool | ../ |
| GITHUB_NAME | github auth credential to push to remote repo | user |
| GITHUB_PASS | github auth credential to push to remote repo | p4ss |

To set these in your environment, you can use direnv (`eval "$(direnv hook bash)"` in Windows Subsystem for Linux), or just export the variables 


```
export USERNAME=test
```

*NOTE: be careful to ignore your .envrc file in .gitignore if you put secrets in it.*