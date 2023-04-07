const faker = require("faker");
const fs = require("fs");
const { Client } = require("pg");

// Your PostgreSQL credentials
const POSTGRES_USER = "postgres";
const POSTGRES_PASSWORD = "password";
const POSTGRES_DB = "mydb";
const POSTGRES_PORT = 5432;
const POSTGRES_HOST = "localhost";

const numUsers = 100;
const numSpecies = 50;
const numAdvices = 150;
const numLocations = 100;
const numPlants = 200;
const numPhotos = 300;
const numKeeps = 100;
const numMessages = 500;
const numNeedKeeps = 50;

function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

function generateUsers() {
  const users = [];
  for (let i = 1; i <= numUsers; i++) {
    users.push({
      id: i,
      username: faker.internet.userName(),
      firstname: faker.name.firstName(),
      lastname: faker.name.lastName(),
      email: faker.internet.email(),
    });
  }
  return users;
}

function generateSpecies() {
  const species = [];
  for (let i = 1; i <= numSpecies; i++) {
    species.push({
      id: i,
      name: faker.random.word(),
    });
  }
  return species;
}

function generateAdvice(species) {
  const advices = [];
  for (let i = 1; i <= numAdvices; i++) {
    advices.push({
      id: i,
      text: faker.lorem.sentence(),
      species_id: randomInt(1, species.length),
    });
  }
  return advices;
}

function generateLocations() {
  const locations = [];
  for (let i = 1; i <= numLocations; i++) {
    locations.push({
      id: i,
      address: faker.address.streetAddress(),
      zip: faker.address.zipCode().substring(0, 5),
      city: faker.address.city(),
      lat: parseFloat(parseFloat(faker.address.latitude()).toFixed(2)),
      lon: parseFloat(parseFloat(faker.address.longitude()).toFixed(2)),
    });
  }
  return locations;
}


function generatePlants(users, species, locations) {
  const plants = [];
  for (let i = 1; i <= numPlants; i++) {
    plants.push({
      id: i,
      location_id: randomInt(1, locations.length),
      species_id: randomInt(1, species.length),
      user_id: randomInt(1, users.length),
    });
  }
  return plants;
}

function generatePhotos(plants) {
  const photos = [];
  for (let i = 1; i <= numPhotos; i++) {
    photos.push({
      id: i,
      url: faker.image.imageUrl(),
      plant_id: randomInt(1, plants.length),
    });
  }
  return photos;
}

function generateKeeps(users, plants, locations) {
  const keeps = [];
  for (let i = 1; i <= numKeeps; i++) {
    keeps.push({
      id: i,
      start_date: faker.date.recent(180),
      end_date: faker.date.future(0.5),
      plant_id: randomInt(1, plants.length),
      user_id: randomInt(1, users.length),
      location_id: randomInt(1, locations.length),
    });
  }
  return keeps;
}

function generateMessages(users) {
  const messages = [];
  for (let i = 1; i <= numMessages; i++) {
    const sender_id = randomInt(1, users.length);
    let receiver_id;
    do {
      receiver_id = randomInt(1, users.length);
    } while (receiver_id === sender_id);
    messages.push({
      id: i,
      sender_id: sender_id,
      receiver_id: receiver_id,
      text: faker.lorem.words(8).substring(0, 50),
    });
  }
  return messages;
}

function generateNeedKeeps(users, plants) {
  const needKeeps = [];
  const uniqueUserIds = new Set();

  while (uniqueUserIds.size < numNeedKeeps) {
    uniqueUserIds.add(randomInt(1, users.length));
  }

  let i = 1;
  for (const user_id of uniqueUserIds) {
    needKeeps.push({
      user_id: user_id,
      start_date: faker.date.recent(180),
      end_date: faker.date.future(0.5),
      plant_id: randomInt(1, plants.length),
    });
    i++;
  }
  return needKeeps;
}


function generateData() {
  const users = generateUsers();
  const species = generateSpecies();
  const advices = generateAdvice(species);
  const locations = generateLocations();
  const plants = generatePlants(users, species, locations);
  const photos = generatePhotos(plants);
  const keeps = generateKeeps(users, plants, locations);
  const messages = generateMessages(users);
  const needKeeps = generateNeedKeeps(users, plants);

  return {
    users,
    species,
    advices,
    locations,
    plants,
    photos,
    keeps,
    messages,
    needKeeps,
  };
}

const data = generateData();
fs.writeFileSync("data.json", JSON.stringify(data, null, 2));

console.log("Data generated successfully in data.json");

async function insertDataToDatabase(data) {
  const client = new Client({
    user: POSTGRES_USER,
    password: POSTGRES_PASSWORD,
    database: POSTGRES_DB,
    port: POSTGRES_PORT,
    host: POSTGRES_HOST,
  });

  try {
    await client.connect();

    for (const user of data.users) {
      await client.query(
        "INSERT INTO users(id, username, firstname, lastname, email) VALUES($1, $2, $3, $4, $5)",
        [user.id, user.username, user.firstname, user.lastname, user.email]
      );
    }

    for (const species of data.species) {
      await client.query("INSERT INTO species(id, name) VALUES($1, $2)", [
        species.id,
        species.name,
      ]);
    }

    for (const advice of data.advices) {
      await client.query(
        "INSERT INTO advice(id, text, species_id) VALUES($1, $2, $3)",
        [advice.id, advice.text, advice.species_id]
      );
    }

    for (const location of data.locations) {
      await client.query(
        "INSERT INTO location(id, address, zip, city, lat, lon) VALUES($1, $2, $3, $4, $5, $6)",
        [
          location.id,
          location.address,
          location.zip,
          location.city,
          location.lat,
          location.lon,
        ]
      );
    }

    for (const plant of data.plants) {
      await client.query(
        "INSERT INTO plants(id, location_id, species_id, user_id) VALUES($1, $2, $3, $4)",
        [plant.id, plant.location_id, plant.species_id, plant.user_id]
      );
    }

    for (const photo of data.photos) {
      await client.query(
        "INSERT INTO photo(id, url, plant_id) VALUES($1, $2, $3)",
        [photo.id, photo.url, photo.plant_id]
      );
    }

    for (const keep of data.keeps) {
      await client.query(
        "INSERT INTO keeps(id, start_date, end_date, plant_id, user_id, location_id) VALUES($1, $2, $3, $4, $5, $6)",
        [
          keep.id,
          keep.start_date,
          keep.end_date,
          keep.plant_id,
          keep.user_id,
          keep.location_id,
        ]
      );
    }

    for (const message of data.messages) {
      await client.query(
        "INSERT INTO message(id, sender_id, receiver_id, text) VALUES($1, $2, $3, $4)",
        [message.id, message.sender_id, message.receiver_id, message.text]
      );
    }

    for (const needKeep of data.needKeeps) {
      await client.query(
        "INSERT INTO need_keep(user_id, start_date, end_date, plant_id) VALUES($1, $2, $3, $4)",
        [
          needKeep.user_id,
          needKeep.start_date,
          needKeep.end_date,
          needKeep.plant_id,
        ]
      );
    }

    console.log("Data inserted successfully into the database");
  } catch (err) {
    console.error("Error inserting data into the database:", err);
  } finally {
    await client.end();
  }
}

(async () => {
  const data = generateData();
  fs.writeFileSync("data.json", JSON.stringify(data, null, 2));
  console.log("Data generated successfully in data.json");

  await insertDataToDatabase(data);
})();
