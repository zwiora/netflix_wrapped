const express = require('express');
const session = require('express-session');
// const data = require('../data/mydata');
const router = express.Router();

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
