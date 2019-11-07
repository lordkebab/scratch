const express = require('express')
app = express()
port = 3000

app.get('/', (req, res) => {
	res.send('Got a GET request at /');
});

app.get('/user', (req, res) => {
	res.send('Got a GET request at /user');
});

app.listen(port);
