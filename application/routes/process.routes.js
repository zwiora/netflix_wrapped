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
  const username = req.session.username;
  const startDate = Date.parse(req.session.startDate);
  const endDate = Date.parse(req.session.endDate);
  // Convert file to JSON
  let activity = {
    profiles: []
  };
  const uploadedFilePath = path.join(__dirname, '../../uploads', req.session.reportId);
  if (!fs.existsSync(uploadedFilePath)) {
    res.redirect('/');
    return;
  }
  const fileStream = fs.createReadStream(uploadedFilePath);
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
    if (profile != username)
      continue;
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
    const currentDate = Date.parse(l[1]);
    if (currentDate < startDate || currentDate > endDate) {
      continue;
    }
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
  
  // Delete uploaded file
  fs.unlink(uploadedFilePath, (_) => {});
  // Check if input is valid
  if (activity.profiles.length == 0)
    return res.status(500).send('No profiles with this name');
  if (activity.profiles[0].viewingActivity.length[0])
    return res.status(500).send('No viewing activity for this profile in this period');


//   // Load test data instead of calling the API
// const testReportPath = path.join(__dirname, '../../API/test_data/example_report.json');

// try {
//   const testReport = JSON.parse(fs.readFileSync(testReportPath, 'utf8'));
//   req.session.report = testReport;

//   fs.writeFile(
//     path.join(__dirname, '../../reports', req.session.reportId),
//     JSON.stringify(testReport),
//     (err) => {
//       if (err) return res.status(500).send('Server failed');
//       return res.status(200).send('OK');
//     }
//   );
// } catch (err) {
//   console.error('Failed to load test report:', err);
//   return res.status(500).send('Failed to load test report');
// }

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
          }
        );
        return res.status(200).send('OK');
      }
    })
    .catch(function (error) {
      console.log(error);
      return res.status(500).send('API failed');
    });
});

module.exports = router;