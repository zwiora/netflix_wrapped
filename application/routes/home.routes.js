const express = require('express');
const _ = require('express-session');
const router = express.Router();
const formidable = require('formidable');
const path = require('path');
const fs = require('fs');
const readline = require('readline');


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

router.post('/process', async (req, res) => {
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
  
  // return res.status(200).send('OK');
  return res.status(500).send('API failed');
});

router.get('/report', (req, res) => {
  const report = req.app.locals.reportData;

  if (!report) {
    return res.status(404).send('No report available. Please upload a file first.');
  }

  res.render('report', { report });
});
// router.get('/getCategories', (req, res) => {
//     res.redirect('/');
// })

// router.get('/getProducts/:id([0-9]{1,2})', (req, res) => {
//     let allProducts = data.categories;
//     let categoryId = parseInt(req.params.id);
//     let thisCategoryProducts = [];
//     let allCategories = [];

//     for (let index = 0; index < allProducts[categoryId].products.length; index++) {
//         const element = allProducts[categoryId].products[index];
//         thisCategoryProducts.push({
//             "name": element.name,
//             "image": element.image,
//             "amount": 0,
//             "id": categoryId + "_" + index
//         });
//     }

//     let cartContentCounter = 0;
//     if (req.session.cart != undefined) {
//         req.session.cart.forEach(cartElement => {
//             cartContentCounter += cartElement.count;
//             if (cartElement.category == categoryId) {
//                 for (let index = 0; index < thisCategoryProducts.length; index++) {
//                     const element = thisCategoryProducts[index];
//                     if (element.id == categoryId + "_" + cartElement.item) {
//                         element.amount = cartElement.count;
//                     }
//                 }
//             }
//         });
//     }

//     for (let index = 0; index < allProducts.length; index++) {
//         const element = allProducts[index];
//         allCategories.push({
//             "name": element.name,
//             "id": index
//         });        
//     }

//     res.render('category', {
//         title: allProducts[categoryId].name,
//         currentCategoryName: allProducts[categoryId].name,
//         cartCounter: cartContentCounter,

//         availableProducts: thisCategoryProducts,
//         categoriesNames: allCategories,
//         currentCategoryId: categoryId
//     });
// });

// router.get('/addToCart/:id([0-9]{1,2}_[0-9]{1,9})', (req, res) => {
//     let complexId = req.params.id.split('_');
//     let categoryId = complexId[0];
//     let productId = complexId[1];

//     if (req.session.cart === undefined) {
//         req.session.cart = [
//             {
//                 "category": categoryId,
//                 "item": productId,
//                 "count": 1
//             }
//         ];
//     } else {
//         let alreadyExists = false;
//         for (let index = 0; index < req.session.cart.length; index++) {
//             const element = req.session.cart[index];
//             if (element.category == categoryId && element.item == productId) {
//                 element.count++;
//                 alreadyExists = true;
//                 break;
//             }
//         }
//         if (!alreadyExists) {
//             req.session.cart.push({
//                 "category": categoryId,
//                 "item": productId,
//                 "count": 1
//             });
//         }
//     }

//     res.redirect('/getProducts/' + categoryId);
// });

module.exports = router;
