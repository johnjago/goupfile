# Changelog

## 0.4.1 (2021-10-31)

### Bug fixes
- Provide a way to specify the host URL in the QR code (don't know how I missed
  this!)

## 0.4.0 (2021-09-21)

### New features
- QR code on the file view page
- File names are shown in the page title for easy identification

### Improvements
- Resize correctly on mobile screens
- Remove lowercase "j" from possible file ID chars because it's too close to an
  "i" in some fonts
- The temporary statement to drop the table upon restart has been removed, so
  now it can be used in production!

### Bug fixes
- Use external hostname for the download link when hosted on localhost behind
  a web server

## 0.3.0 (2021-07-22)

### New features
- Support for more than one file
- Button to go back home from the file view page
- Show a message if a file doesn't exist

### Improvements
- Better HTTP request logging
- Remove ambiguous characters from file IDs
- Make the download links actually look like links

## 0.2.0 (2021-06-23)

### New features
- There's now a web UI for uploading files! You can try it at https://goupfile.com

### Improvements
- Remove a lot of unnecessary code and dependencies
- Now styled with Tailwind CSS for a more consistent design
- Easier to deploy since it now uses SQLite as the default database instead of MariaDB

## 0.1.0 (2019-08-18)

### New features
- This is the initial release which provides HTTP API endpoints for uploading
  a file and downloading a file by ID
