const iconeSearch = document.getElementById('iconeSearch');
const iconeFilter = document.getElementById('iconeFilter');
const formSearch = document.getElementById('searchForm');
const formFilter = document.getElementById('filterForm');


iconeSearch.addEventListener('click', () => {
    if (formSearch.style.display === 'none') {
        formSearch.style.display = 'flex';
        formFilter.style.display = 'none';



    } else { formSearch.style.display = 'none'; }
});
iconeFilter.addEventListener('click', () => {
    if (formFilter.style.display === 'none') {
        formFilter.style.display = 'flex';
        // form2.style.alignItems = 'center';
        // form2.style.justifyContent='center';
        // form2.style.marginTop='15px';

        formSearch.style.display = 'none';


    } else { formFilter.style.display = 'none'; }
});

// // const icone = document.getElementById('icone');
// //  const conteneur = document.getElementById('container');

// // icone.addEventListener('click', () => { conteneur.classList.toggle('caché'); conteneur.classList.toggle('visible'); });

// function toggleFormulaire() { var formulaire = document.getElementById("formulaire"); 
// formulaire.classList.toggle("affiche"); } 
// const first = document.getElementById('first');
// const second = document.getElementById('second');
// const testf = document.getElementById('sliderOutput1');

// https://upload.wikimedia.org/wikipedia/commons/thumb/2/2c/Rotating_earth_%28large%29.gif/250px-Rotating_earth_%28large%29.gif

// window.addEventListener("DOMContentLoaded", (event) => {
//     // console.log("DOM entièrement chargé et analysé");
//     testf.innerHTML="1234"
//     console.log(testf)
//   });


