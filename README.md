# Endito
A service for editing web pages.

## Support
If you enjoy the service and want to support me in further open source development.

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=8P7E4L2D2K6BU)

## Content editing
Any content editing system must have a few components
- awareness of the source files that need to be updated
- input system to receive file update requests
- a translation of input to file schema or syntax

## Existing solutions
This is not a new problem and many solutions have been developed with several more features than listed above. Some important additional features
- authentication/authorization to ensure edits are made by allowed users
- user interface that allows users to easily, confidently make updates
- version control for files to revert 

## Dependencies
- File System (directory) with read and write access
- Shared syntax between the files and the input requests

## How it works
Any HTML file can be fetched and Endito renders it with a WYSIWYG-style wrapper. After edits are performed through Endito's editor, the editor can enter their username and password to submit the edit request. If the HTML file's host server supports Endito, and the editor's authorization is successful, the files will be uploaded and the new HTML files will be present.

## Usage
Endito is run as a companion service on the same host as the HTML file server. It accepts a path to the HTML file root directory on startup and can fetch auth from a variety of ways. It receives a POST request from the user interface (WYSIWYG editor), performs a git commit, and updates the POST'd file.

### Notes
WSL direnv
- eval "$(direnv hook bash)"