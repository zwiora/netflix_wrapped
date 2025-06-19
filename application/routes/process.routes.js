const express = require('express');
const _ = require('express-session');
const router = express.Router();
const path = require('path');
const fs = require('fs');
const readline = require('readline');
const http = require('http');


router.post('/', async (req, res) => {
  if (!req.session.reportId) {
    res.redirect('/');
  }
  // Convert file to JSON
  let activity = {
    profiles: []
  };
  const fileStream = fs.createReadStream(path.join(__dirname, '../../uploads', req.session.reportId));
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity
  });
  firstLine = true;
  for await (const line of rl) {
    if (firstLine) {
      firstLine = false;
      continue;
    }
    // Split line using ',' (except when in quotes)
    let l = line.replaceAll(",,", ', ,');
    l = Array.from(l.matchAll(/[^",]+|"([^"]*)"/g), ([a,b]) => b || a);
    // Check if profile already exists on activity list, set profile variable to its ID
    let profile = l[0];
    let profileExists = false;
    for (let index = 0; index < activity.profiles.length; index++) {
      const element = activity.profiles[index];
      if (element.name == profile) {
        profileExists = true;
        profile = index;
        break;
      }
    }
    if (!profileExists) {
      activity.profiles.push({
        name: profile,
        viewingActivity: []
      });
      profile = activity.profiles.length - 1;
    }
    // Add current line to the viewingActivity
    activity.profiles[profile].viewingActivity.push({
      startTime: l[1],
      duration: l[2],
      attributes: l[3],
      title: l[4],
      deviceType: l[6],
      bookmark: l[7],
      latestBookmark: l[8],
      country: l[9]
    })
  }
  // Send request to API
  // let activity = {"profiles":[{"name":"Łukasz","viewingActivity":[{"startTime":"2020-05-06 09:20:01","duration":"00:21:46","attributes":"Brooklyn 9-9: Sezon 2: Tajniak (Odcinek 1)","title":"Netflix Opera Other","deviceType":"00:21:43","bookmark":"PL (Poland)"},{"startTime":"2020-05-05 21:50:32","duration":"00:21:39","attributes":"Brooklyn 9-9: Sezon 1: Zmiany, zmiany (Odcinek 22)","title":"Android DefaultWidevineL3Phone Android Phone","deviceType":"00:21:35","bookmark":"PL (Poland)"},{"startTime":"2020-05-05 21:27:01","duration":"00:21:36","attributes":"Brooklyn 9-9: Sezon 1: Nie do rozwiązania (Odcinek 21)","title":"Android DefaultWidevineL3Phone Android Phone","deviceType":"00:21:36","bookmark":"PL (Poland)"}]}]};
  
  const options = {
    hostname: 'localhost',
    port: 8080,
    path: '/generate',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'User-Agent': 'Node.js'
    }
  };
  const activityString = JSON.stringify(activity);
  options.headers['Content-Length'] = activityString.length;
  const request = http.request(options, (response) => {
    let analyzedData = '';

    response.on('data', (chunk) => {
      analyzedData += chunk.toString();
    });
    response.on('end', () => {
      // const result = JSON.parse(analyzedData);
      console.log(analyzedData);

      console.log(`Status code: ${response.statusCode}`);
      console.log(`Headers: ${JSON.stringify(response.headers)}`);
      // console.log(`Post ID: ${result.id}`);
      // console.log(`Post Title: ${result.title}`);
      // console.log(`Post Body: ${result.body}`);
      // console.log(`Post User ID: ${result.userId}`);

      if (response.statusCode == 200)
        return res.status(200).send('OK');
      else
        return res.status(500).send('API failed');
    })
    response.on('error', (error) => {
      return res.status(500).send('API failed');
    })    
  });

  request.write(activityString);
  request.end();  
});

module.exports = router;
