import pymongo

# create a new MongoClient instance
client = pymongo.MongoClient("mongodb://localhost:27017/")

# switch to the database that contains the cached movie data
db = client["movie_cache"]

# query the movies collection to check if the data is present
movies = db.movies.find()

# print the movie data
for movie in movies:
    print(movie)
