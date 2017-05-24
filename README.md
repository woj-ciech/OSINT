# OSINT

Small scripts for OSINT.<br>
#1. Instagram<br>
Description:<br>
Small proof of concept to show how to retrieve exact location of photos from instagram. At instagram.com you can only see approximate location like country and city. This script gets latitude and longitude from photos and next checks this values with GeoPy. Additionaly can count all hashtags.<br>
<h2>Tested with >600 photos<br>
Do not forget to change api credentials</h2>

Functions:
- Count and show all hashtags
- Get exact location of all photos (thanks to GeoPy) with timestamp 

Requirements:
- GeoPy (pip install geopy)
- InstagramApi for python (https://github.com/LevPasha/Instagram-API-python) (pip install -e git+https://github.com/LevPasha/Instagram-API-python.git#egg=InstagramAPI)

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
