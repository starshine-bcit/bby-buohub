const express = require('express');
const mysql = require("mysql");
const dotenv = require('dotenv');

const app = express();
//const bcrypt = require("bcryptjs")


//dotenv.config({ path: './.env' });

const db = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'ACIT3855',
    database: 'logindb'
});

/* const db = mysql.createConnection({
    host: process.env.HOST,
    user: process.env.USER,
    password: process.env.PASSWORD,
    database: process.env.DATABASE
}); 
*/

// Connect to MySQL database
db.connect((err) => {
    if (err) {
        console.error('Error connecting to database:', err);
        return;
    }
    console.log('Connected to the database');
});
 

app.set('view engine', 'hbs')


const path = require('path');

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

app.post("/auth/register", (req, res) => {    
    const { username, password, passwordConfirm } = req.body;
    //console.log(req.body)
      db.query('SELECT username FROM users WHERE username = ?', [username], async (error, result) => {
        if(error) {
            console.log(error);
        }
        if(result.length > 0) {
            return res.render('register', {
                message: 'This username is already in use'
            });
        } 
        
        if( password !== passwordConfirm) {
            return res.render('register', {
                message: 'Passwords do not match!'
            });
        }
        
       // let hashedPassword = await bcrypt.hash(password, 8);

        db.query('INSERT INTO users SET ?', { username: username, password: password }, (err, result) => {
            if(err) {
                console.log(err);
            } else {
                return res.render('register', {
                    message: 'User registered!'
                });
            }
        });
    });
});


app.listen(5000, ()=> {
    console.log("Server started on port 5000")
})