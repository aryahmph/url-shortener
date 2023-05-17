db.createUser({
    user: "aryahmph",
    pwd: "aryahmph",
    roles: [
        {
            role: "readWrite",
            db: "url_shortener",
        },
    ],
});

db = new Mongo().getDB("url_shortener");
db.createCollection("links");
