const express = require('express');
const mysql = require("mysql");
require('dotenv').config();
var favicon = require('serve-favicon')
const path = require('path');
var cookieParser = require('cookie-parser');
const axios = require('axios')

const app = express();
const authBaseURL = `http://${process.env.AUTH_HOST}:${process.env.AUTH_PORT}`
const cdnBaseURL = `http://${process.env.CDN_HOST}:${process.env.CDN_PORT}`


const db = mysql.createConnection({
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  port: parseInt(process.env.DB_PORT)
});


db.connect((err) => {
  if (err) {
    console.error('Error connecting to MariaDB:', err);
    return;
  }
  console.log('Connected to the MariaDB');
});

app.set('view engine', 'hbs')
app.use(cookieParser());
app.use(favicon(path.join(__dirname, './public', 'favicon.ico')))
const publicDir = path.join(__dirname, './public');
app.use(express.static(publicDir));

app.get("/register", (req, res) => {
  res.render("register")
})

app.get("/login", (req, res) => {
  res.render("login")
})

app.get("/ping", (req, res) => {
  res.send("pong")
})

app.use((req, res, next) => {
  res.setHeader('Access-Control-Allow-Origin', 'http://localhost:9001');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
  res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
  res.setHeader('Access-Control-Allow-Credentials', true);
  const url = new URL(req.url, `http://${req.headers.host}`)
  if (url.pathname === '/login' || url.pathname === '/register' || url.pathname === '/auth/login' || url.pathname === '/auth/register' || url.pathname === '/ping') {
    next();
    return
  }
  const authCookie = req.cookies.AuthToken;
  const refreshCookie = req.cookies.RefreshToken;
  if (authCookie === undefined || refreshCookie === undefined) {
    res.redirect('/login');
    return;
  }
  axios.post(`${authBaseURL}/auth`, {
    accessToken: authCookie,
    refreshToken: refreshCookie
  })
    .then((response) => {
      if (response.status !== 202 || response.data.ok !== true) {
        console.log(response);
        res.redirect('/login');
        return;
      }
      if (response.data.newToken !== "") {
        res.cookie('AuthToken', response.data.newToken, {
          maxAge: 60 * 60 * 24,
          path: '/'
        })
      }
      next();
      return;
    })
    .catch((error) => {
      console.log(error);
      res.redirect('/login');
      return;
    })
});

app.use(express.urlencoded({ extended: false }));
app.use(express.json());

app.post('/auth/register', (req, res) => {
  axios.post(`${authBaseURL}/create`, req.body)
    .then((response) => {
      if (response.status !== 201 || response.data.created !== true) {
        console.log(response);
        res.redirect('/register');
        return;
      }
      res.cookie('AuthToken', response.data.accessToken, {
        maxAge: 60 * 60 * 24,
        path: '/'
      })
      res.cookie('RefreshToken', response.data.refreshToken, {
        maxAge: 60 * 60 * 24,
        path: '/'
      })
      res.redirect('/video')
    })
    .catch((error) => {
      console.log(error);
      res.redirect('/register');
      return;
    })
});

app.post('/auth/login', (req, res) => {
  axios.post(`${authBaseURL}/login`, req.body)
    .then((response) => {
      if (response.status !== 202 || response.data.valid !== true) {
        console.log(response);
        res.redirect('/login');
        return;
      }
      res.cookie('AuthToken', response.data.accessToken, {
        maxAge: 60 * 60 * 24,
        path: '/'
      })
      res.cookie('RefreshToken', response.data.refreshToken, {
        maxAge: 60 * 60 * 24,
        path: '/'
      })
      res.redirect('/video')
    })
    .catch((error) => {
      console.log(error);
      res.redirect('/login');
      return;
    })
});



app.get('/play/:uuid/:manifest_name', (req, res) => {
  const uuid = req.params.uuid;
  const manifestName = req.params.manifest_name;

  res.redirect(`/player/${uuid}/${manifestName}`);
});

app.get('/player/:uuid/:manifest_name', (req, res) => {
  const uuid = req.params.uuid;
  const manifestName = req.params.manifest_name;
  const videoUrl = {
    url: `${cdnBaseURL}/stream/${uuid}/${manifestName}`,
    poster: `${cdnBaseURL}/stream/${uuid}/thumb.png`
  }
  res.render('player', videoUrl);
});

app.get('/video', (req, res) => {
  const videoQuery = 'SELECT * FROM videos WHERE process_complete = true ORDER BY id DESC LIMIT 9';
  db.query(videoQuery, (videoError, videoResults) => {
    if (videoError) {
      console.error('Error executing video query:', videoError);
      res.status(500).send('Internal Server Error');
      return;
    }
    const imagesArray = videoResults.map(video => ({ uuid: video.uuid, poster_filename: video.poster_filename, title: video.title, description: video.description, manifest_name: video.manifest_name }));
    const urlArray = imagesArray.map(image => ({
      url: `${cdnBaseURL}/stream/${image.uuid}/${image.poster_filename}`,
      title: image.title,
      description: image.description,
      videoUrl: `${cdnBaseURL}/stream/${image.uuid}/${image.manifest_name}`,
      uuid: image.uuid,
      manifest_name: image.manifest_name
    }));
    res.render('video', { urlArray: urlArray });
  });
});

app.get("/", (req, res) => {
  res.render("index")
})

app.listen(8999, () => {
  console.log("Server started on port 8999")
})

