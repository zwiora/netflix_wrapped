const express = require('express');
const session = require('express-session');
// const data = require('../data/mydata');
const axios = require('axios');
const router = express.Router();
const fs = require('fs');
const { IncomingForm } = require('formidable');
const path = require('path');
const uploadDir = path.join(__dirname, '../../uploads');
const FormData = require('form-data');
if (!fs.existsSync(uploadDir)) {
  fs.mkdirSync(uploadDir, { recursive: true });
}

const form = new IncomingForm({
  uploadDir,
  keepExtensions: true,
});

router.get('/', function(req, res, next) {
    // allCategories = [];
    // for (let index = 0; index < data.categories.length; index++) {
    //     const element = data.categories[index];
    //     allCategories.push({
    //         "name": element.name,
    //         "id": index
    //     });        
    // }

    // let cartContentCounter = 0;
    // if (req.session.cart != undefined) {
    //     req.session.cart.forEach(cartElement => {
    //         cartContentCounter += cartElement.count;
    //     });
    // }

    res.render('home', {
        title: 'Netflix Wrapped',
        // currentCategoryName: 'Welcome!',
        // cartCounter: cartContentCounter,
        // categoriesList: allCategories
    });
});

router.post('/upload', (req, res) => {
  // Just save file temporarily, then redirect to waiting screen
  const form = new IncomingForm({
    uploadDir: path.join(__dirname, '../../uploads'),
    keepExtensions: true,
  });

  form.parse(req, async (err, fields, files) => {
    if (err || !files.netflixData) {
      return res.status(400).send('No file uploaded.');
    }

    const file = Array.isArray(files.netflixData) ? files.netflixData[0] : files.netflixData;

    // Save filename in app memory (or session, if preferred)
    req.app.locals.pendingUpload = file;
    res.redirect('/waiting');
  });
});
router.post('/process', async (req, res) => {
  const file = req.app.locals.pendingUpload;

  if (!file) {
    return res.status(400).send('No pending file');
  }

  const fileStream = fs.createReadStream(file.filepath);
  const formData = new FormData();
  formData.append('file', fileStream, file.originalFilename);

  try {
    const apiResponse = await axios.post('http://localhost:5000/analyze', formData, {
      headers: formData.getHeaders(),
      timeout: 4000, // Optional: shorten request timeout
    });

    req.app.locals.reportData = apiResponse.data;
    return res.status(200).send('OK');
  } catch (err) {
    console.error('API call failed:', err.message);
    return res.status(500).send('API failed');
  }
});
router.get('/waiting', (req, res) => {
  res.render('waiting');
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
