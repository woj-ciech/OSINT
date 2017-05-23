# OSINT

Small scripts for OSINT.
#1. Instagram
Description
Small proof of concept to show how to retrieve exact location of photos from instagram. At instagram.com you can only see Approximate location like country and city. This script gets latitude and longitude from photos with location and next check this values with GeoPy. Additionaly can count all hashtags (you can add it to your dictionary).

Function:
- Count and present hashtags
- Get exact location of all photos (thanks to GeoPy)

Usage:
./script "ID"
You can check user ID easily by adding ?_a=1 to the end of the request, like this https://www.instagram.com/USERNAME/?__a=1 and look for id on json response.

Examples:
