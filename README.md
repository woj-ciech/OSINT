# OSINT

Small scripts for OSINT.<br>
#1. Instagram<br>
Description:<br>
Small proof of concept to show how to retrieve exact location of photos from instagram. At instagram.com you can only see approximate location like country and city. This script gets latitude and longitude from photos and next checks this values with GeoPy. Additionaly can count all hashtags.

Functions:
- Count and show all hashtags
- Get exact location of all photos (thanks to GeoPy) with timestamp 

Requirements:
- GeoPy 
- InstagramApi for python (https://github.com/LevPasha/Instagram-API-python)

Usage:
#python insta.py ID<br>
You can check user's ID easily by adding double underscore at the end of the request, like this https://www.instagram.com/USERNAME/?__a=1 and look for owner id in json response.

Examples: (random person)

![alt text](http://i.imgur.com/2eOwovn.png)

![alt text](http://i.imgur.com/WTVkFPM.png)
TODO:
- improve hashtags parsing
- make it faster
- new functions?

Code is quite messy but in my defense I can say that I'm still learning python.
If you have any advice or idea, just let me know!
