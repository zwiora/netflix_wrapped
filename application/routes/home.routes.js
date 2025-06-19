const express = require('express');
const _ = require('express-session');
const router = express.Router();
const formidable = require('formidable');
const path = require('path');
const fs = require('fs');


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
  const form = new formidable.IncomingForm({
    uploadDir: path.join(__dirname, '../../uploads'),
    keepExtensions: true
  });

  form.parse(req, function (err, fields, files) {
    if (err || !files.netflixData) {
      return res.status(400).json({ success: false, error: 'No file uploaded.' });
    }

    const uploadedFile = Array.isArray(files.netflixData)
      ? files.netflixData[0]
      : files.netflixData;

    // Save it temporarily (could store path in session, too)
    req.session.uploadedFile = uploadedFile;

    return res.status(200).json({
      success: true,
      filename: uploadedFile.originalFilename
    });
  });
});


router.get('/report', (req, res) => {
  if (!req.session.reportId) {
    res.redirect('/');
    return;
  }
  res.render('report', {
    title: "Netflix Wrapped - report",
    report: req.session.report,
    reportId: req.session.reportId
  });
});

router.get('/report/:id', (req, res) => {
  try {
    req.session.reportId = req.params.id;
    req.session.report = JSON.parse(fs.readFileSync(path.join(__dirname, '../../reports', req.session.reportId), 'utf8'));
    res.redirect('/report');
  } catch (_) {
    res.redirect('/');
  }
})

module.exports = router;
