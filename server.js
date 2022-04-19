import xp from 'express';

const app = xp();
const PORT = 8080;

app.use(xp.static('public'));

app.get('/', (req, res) => {
    res.send('Hussain');
});

app.listen(PORT, () => console.log(`Server running at port ${PORT}`));
