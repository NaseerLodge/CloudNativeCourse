
from flask import Flask, render_template, request
from movieclient import movieapi
import requests
import pymongo


# create a new MongoClient instance
client = pymongo.MongoClient("mongodb://localhost:27017/")

# create a new database and collection
db = client["movie_cache"]
collection = db["movies"]

api_key = "4e57a63"

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def home():
    if request.method == 'POST':
        st = request.form['movie_name']
        n = request.form['no_of_recommendations']
        movies = movieapi(st, int(n))
        print(movies)
        movies = movie_info(movies)

        return render_template('recommendations.html', movies=movies)
    else:
        return render_template('home.html')

@app.route('/recommendations')
def recommendations():
    movies = request.args.getlist('movies')
    return render_template('recommendations.html', movies=movies)

'''
@app.route('/recommendations')
def movie_info(movies):
    movie_details = []
    for movie_name in movies:
        # Construct the API request URL with the movie title parameter
        url = f"http://www.omdbapi.com/?apikey={api_key}&t={movie_name}"

        # Send the API request and retrieve the response data
        response = requests.get(url)
        data = response.json()

        # Extract the year, released date, genre, language, poster and type from the response data
        year = data["Year"]
        if '–' in year:
            year = year.split('–')[0]
        released = data["Released"]
        genre = data["Genre"]
        language = data["Language"]
        poster = data["Poster"]
        type = data["Type"]

        movie_information = [str(movie_name), int(year), str(released), str(genre), str(language), str(poster), str(type)]
        movie_details.append(movie_information)

    return movie_details

'''

@app.route('/recommendations')
def movie_info(movies):
    movie_details = []
    for movie_name in movies:
        # Check if the movie is present in the database
        movie = collection.find_one({'title': movie_name})
        if movie:
            # If the movie is present in the database, retrieve its details from there
            year = movie['year']
            released = movie['released']
            genre = movie['genre']
            language = movie['language']
            poster = movie['poster']
            type = movie['type']
        else:
            # If the movie is not present in the database, fetch its details from the API and save them in the database
            url = f"http://www.omdbapi.com/?apikey={api_key}&t={movie_name}"
            response = requests.get(url)
            data = response.json()
            year = data.get('Year', 'N/A')
            if year:
                if '–' in year:
                    year = year.split('–')[0]
            released = data.get('Released', 'N/A')
            genre = data.get('Genre', 'N/A')
            language = data.get('Language', 'N/A')
            poster = data.get('Poster', 'N/A')
            type = data.get('Type', 'N/A')
            # Save the movie details in the database
            movie_dict = {
                'title': movie_name,
                'year': str(year),
                'released': str(released),
                'genre': str(genre),
                'language': str(language),
                'poster': str(poster),
                'type': str(type)
            }
            collection.insert_one(movie_dict)

        movie_information = [str(movie_name), str(year), str(released), str(genre), str(language), str(poster), str(type)]
        movie_details.append(movie_information)

    return movie_details




if __name__ == '__main__':
    app.run(debug=True)
