const express = require('express');
const _ = require('express-session');
const router = express.Router();
const formidable = require('formidable');
const path = require('path');


router.get('/', function(_, res, next) {
  res.render('home', {
    title: 'Netflix Wrapped'
  });
});

router.get('/waiting', (_, res) => {
  res.render('waiting', {
    title: 'Processing data...'
  });
});

router.post('/upload', (req, res) => {
  let form = new formidable.IncomingForm({
    uploadDir: path.join(__dirname, '../../uploads'),
    keepExtensions: true
  });
  form.parse(req, function (err, fields, files) {
    if (err || !files.netflixData) {
      return res.status(400).send('No file uploaded.');
    }
    const uploadedFile = Array.isArray(files.netflixData) ? files.netflixData[0] : files.netflixData;
    req.session.reportId = uploadedFile.newFilename;
    res.redirect('/waiting');
  })
});

router.get('/report', (req, res) => {
  const report = req.app.locals.reportData;

  if (!report) {
    return res.status(404).send('No report available. Please upload a file first.');
  }

  res.render('report', { report });
});

module.exports = router;
