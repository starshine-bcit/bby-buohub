const express = require('express');
const mysql = require("mysql");
const dotenv = require('dotenv');
const http = require('http');
const app = express();
var favicon = require('serve-favicon')
const path = require('path');

//const bcrypt = require("bcryptjs")

app.use(favicon(path.join(__dirname, './public', 'favicon.ico')))


dotenv.config({ path: './.env' });

const db = mysql.createConnection({
    host: 'localhost',
    user: 'auth',
    password: '123test',
    database: 'auth'
});

// const db = mysql.createConnection({
//    host: process.env.HOST,
//    user: process.env.USER,
//    password: process.env.PASSWORD,
//    database: process.env.DATABASE
//}); 


// Connect to MySQL database
db.connect((err) => {
    if (err) {
        console.error('Error connecting to database:', err);
        return;
    }
    console.log('Connected to the database');
});
 

app.set('view engine', 'hbs')




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
    hostname: '127.0.0.1',
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

  // Write the data to the request body and end the request
  request.write(JSON.stringify(dataToSend));
  request.end();
  res.redirect('/login');
});

app.post('/auth/login', (req, res) => {

  // Retrieve the JSON object from the request body
  const dataToSend = req.body;

  // Configure the request details
  const options = {
    hostname: '127.0.0.1',
    port: 9000,
    path: '/login',
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

  // Write the data to the request body and end the request
  request.write(JSON.stringify(dataToSend));
  request.end();
  res.redirect('/');
});
    

app.listen(5000, ()=> {
    console.log("Server started on port 5000")
})