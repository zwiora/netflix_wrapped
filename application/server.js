const express = require('express');
const session = require('express-session');
const app = express();
const port = 3000;
function logger(req, res, next) {
    console.log("Log: " + req.originalUrl);
    next();
}

app.set('view engine', 'ejs');

app.use(session({
    secret: 'secret',
    cookie: { maxAge: 3600000 },
    resave: false,
    saveUninitialized: true
}));
app.use(express.static(__dirname + '/public'));
app.use(express.json())
app.use(logger);

const homeRouter = require('./routes/home.routes');
const processRouter = require('./routes/process.routes');

app.use('/', homeRouter);
app.use('/process', processRouter);

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
  });