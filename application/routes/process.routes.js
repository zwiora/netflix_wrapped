const express = require('express');
const _ = require('express-session');
const router = express.Router();
const path = require('path');
const fs = require('fs');
const readline = require('readline');
const axios = require('axios');


router.post('/', async (req, res) => {
  if (!req.session.reportId) {
    res.redirect('/');
    return;
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
  axios.post('http://localhost:8080/generate', activity)
  .then(function (response) {
    if (response.status != 200) 
      return res.status(500).send('API failed');
    else {
      req.session.report = response.data;
      fs.writeFile(
        path.join(__dirname, '../../reports', req.session.reportId), 
        JSON.stringify(response.data), 
        (err) => {
          if (err)
            return res.status(500).send('Server failed');
          else
            return res.status(200).send('OK');
        }
      );
    }
  })
  .catch(function (error) {
    console.log(error);
    return res.status(500).send('API failed');
  });
});

module.exports = router;
