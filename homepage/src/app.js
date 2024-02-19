const express = require('express');
const mysql = require("mysql");
const dotenv = require('dotenv')
const http = require('http');
const app = express();
var favicon = require('serve-favicon')
const path = require('path');
//var cookieParser = require('cookie-parser')
//const bcrypt = require("bcryptjs")
const axios = require('axios');
dotenv.config();
//const mariadb = require('mariaDB');




//app.set('view engine', 'ejs');


//const exphbs  = require('express-handlebars');

//app.engine('.hbs', exphbs({ extname: '.hbs' }));
//app.set('view engine', '.hbs');

app.use(favicon(path.join(__dirname, './public', 'favicon.ico')))

//app.use('/video', express.static(path.join(__dirname, './public')));



const db = mysql.createConnection({
    host: "db",
    user: 'dbuser',
    password: '123test',
    database: 'bby',
    port: 3306
});

//const pool = mariadb.createPool({
//  host: "db",
//  user: 'dbuser',
//  password: '123test',
//  database: 'bby',
//  connectionLimit: 5
//});

//const db = mysql.createConnection({
//    host: process.env.HOST,
//    user: process.env.USER,
//    password: process.env.PASSWORD,
//    database: process.env.DATABASE
//}); 


// Connect to MySQL database
db.connect((err) => {
    if (err) {
        console.error('Error connecting to MariaDB:', err);
        return;
    }
    console.log('Connected to the MariaDB');
});
 

app.set('view engine', 'hbs')

function showVideo(data) {
  return new Promise((resolve, reject) => {
      axios.get(`http://cdn:9001/stream/${videoData.uuid}/${videoData.manifest_name}`)
          .then(response => {
              console.log('Data sent to auth endpoint successfully.');
              resolve(response.data.video_url); // Assuming the response contains the video URL
          })
          .catch(error => {
              console.error('Error sending data to auth endpoint:', error);
              reject(error);
          });
  });
}

app.get('/video', (req, res) => {
  const cdn = process.env.CDN_HOST
  
  const query = `
    SELECT username
    FROM users
    ORDER BY last_login DESC
    LIMIT 1
  `;


  db.query(query, (error, results) => {
    if (error) {
      console.error('Error executing query:', error);
      res.status(500).send('Internal Server Error');
      return;
    }

    if (results.length === 0) {
      res.status(404).send('No users found');
      return;
    }

    const latestUser = results[0].username;

    // Query to get video information 
    const videoQuery = 'SELECT * FROM videos ORDER BY id DESC LIMIT 9';

    db.query(videoQuery, (videoError, videoResults) => {
      if (videoError) {
        console.error('Error executing video query:', videoError);
        res.status(500).send('Internal Server Error');
        return;
      }

      const imagesArray = videoResults.map(video => ({ uuid: video.uuid, poster_filename: video.poster_filename, title: video.title, description: video.description, manifest_name: video.manifest_name }));

      

      const urlArray = imagesArray.map(image => ({
        url: `http://localhost:9001/stream/${image.uuid}/${image.poster_filename}`,
        title: image.title,
        description: image.description,
        videoUrl: `http://localhost:9001/stream/${image.uuid}/${image.manifest_name}`
    }));
    

      //console.log(urlArray[5].url)

      res.render('video', { latestUser: latestUser, urlArray: urlArray });
    });
  });
});







// Define the directory for serving static files
const publicDir = path.join(__dirname, './public');

// Configure Express to serve static files from the public directory
app.use(express.static(publicDir));

app.get("/", (req, res) => {
    res.render("index")
})

app.get("/register", (req, res) => {
    res.render("register")
})

app.get("/login", (req, res) => {
    res.render("login")
})

app.get("/video", (req, res) => {
  res.render("video")
})

app.use(express.urlencoded({ extended: false }));
app.use(express.json());

//const axios = require('axios');


//app.post("/auth/register", (req, res) => {  
//    const { data }  = req.body 

//    const url = 'http://localhost:9000/create'

// Assuming this is your route handler
//    axios.post(url, data)
//    const { username, password, passwordConfirm } = req.body;
    //console.log(req.body)
  //    db.query('SELECT username FROM users WHERE username = ?', [username], async (error, result) => {
    //    if(error) {
    //        console.log(error);
    //    }
     //   if(result.length > 0) {
      //      return res.render('register', {
      //          message: 'This username is already in use'
     //       });
     //   } 
        
  //      if( password !== passwordConfirm) {
    //        return res.render('register', {
     //           message: 'Passwords do not match!'
    //        });
    //    }
        
       // let hashedPassword = await bcrypt.hash(password, 8);

    //    db.query('INSERT INTO users SET ?', { username: username, password: password }, (err, result) => {
    //        if(err) {
    //            console.log(err);
     //       } else {
     //           return res.render('register', {
      //              message: 'User registered!'
      //          });
    //        }
   //     });
  //  });
//});
app.post('/auth/register', (req, res) => {

  // Retrieve the JSON object from the request body
  const dataToSend = req.body;

  // Configure the request details
  const options = {
    hostname: 'auth',
    port: 9000,
    path: '/create',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    }
  };

  // Create the request
  const request = http.request(options, (response) => {
    let responseData = '';

    // Concatenate chunks of data
    response.on('data', (chunk) => {
      responseData += chunk;
    });

    // Once the response is complete, handle it
    response.on('end', () => {
      console.log('Response from server:', responseData);
      res.send(responseData);
    });
  });

  // Handle errors
  request.on('error', (error) => {
    console.error('Error sending request:', error);
    res.status(500).send('Error sending request');
  });
  //res.redirect('/login');
  //console.log('Cookies: ', req.cookies)
  // Write the data to the request body and end the request
  request.write(JSON.stringify(dataToSend));
  request.end();
  //res.redirect('/login')
});


app.post('/auth/login', (req, res) => {
  const dataToSend = req.body;
  const loginOptions = {
      hostname: 'auth',
      port: 9000,
      path: '/login',
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
      }
  };

  const loginRequest = http.request(loginOptions, (loginResponse) => {
      let responseData = '';
      loginResponse.on('data', (chunk) => {
          responseData += chunk;
      });
      loginResponse.on('end', () => {
          console.log('Response from /auth/login:', responseData);
          sendToAuth(responseData);
          res.redirect('/video');
      });
  });

  // Write the data to the request body for /auth/login and end the request
  loginRequest.write(JSON.stringify(dataToSend));
  loginRequest.end();
});

function sendToAuth(data) {
  axios.post('http://auth:9000/auth', data)
      .then(response => {
          console.log('Data sent to auth endpoint successfully.');
      })
      .catch(error => {
          console.error('Error sending data to auth endpoint:', error);
      });
}

app.listen(8999, ()=> {
    console.log("Server started on port 8999")
})

