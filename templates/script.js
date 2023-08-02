var tbody = document.getElementById('tbody')
var map0 = document.getElementById('map1')
var polyline = []
if (tbody) {
    var tablines = tbody.rows;

    // iterer dans les lignes
    var tabresult = []
    var lentbody = tablines.length
    var i = 0;
    while (i < lentbody) {
        var line = tablines[i]
        var tabcells = line.cells
        // var lenline = tabcells.length
        // console.log(tabligne[0].innerHTML)
        tabresult.push({ city: tabcells[0].innerHTML, latitude: parseFloat(tabcells[1].innerHTML), longitude: parseFloat(tabcells[2].innerHTML) })
        i++;
    }
    // console.log(tab)
    //iterer dans tab
    // console.log(map)
polyline.push([tabresult[0].latitude,tabresult[0].longitude])
    var map = L.map('map').setView([tabresult[0].latitude, tabresult[0].longitude], 13);
    var marker = L.marker([tabresult[0].latitude, tabresult[0].longitude]).addTo(map);
    marker.bindPopup(tabresult[0].city).openPopup();
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 13,
        // noWrap: true,
        attribution: '© OpenStreetMap'
    }).addTo(map);
    console.log("ville", tabresult[0].city, "lar", tabresult[0].latitude, "lon", tabresult[0].longitude)

    var k = 1
    while (k < tabresult.length) {
        console.log("ville", tabresult[k].city, "lar", tabresult[k].latitude, "lon", tabresult[k].longitude)
        var marker = L.marker([tabresult[k].latitude, tabresult[k].longitude]).addTo(map);
        marker.bindPopup(tabresult[k].city).openPopup();
polyline.push([tabresult[k].latitude,tabresult[k].longitude])

        k++;
    }
	

var eskimon = L.polyline(polyline, {color: 'red'}).addTo(map);
console.log(eskimon)

}
if (map0) {
    console.log(map0)
    // 16.0572 et -16.4556 
    var map = L.map('map1').setView([16.0572,-16.4556], 13);
    var marker = L.marker([16.0572,-16.4556]).addTo(map);
    marker.bindPopup("Dakar").openPopup();
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 13,
        attribution: '© OpenStreetMap'
    }).addTo(map);
}


