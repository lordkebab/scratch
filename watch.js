// watch /etc/hosts for changes
const fs = require('fs')

fs.watch('/etc/hosts', (eventType, filename) => {
	console.log(eventType);
	console.log(filename);
});
