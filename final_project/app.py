from flask import Flask, request, render_template
import movieclient

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/recommendations', methods=['POST'])
def recommendations():
    moviename = request.form['moviename']
    no_of_recommendations = int(request.form['no_of_recommendations'])
    recommended_movies = movieclient.movieapi(moviename, no_of_recommendations)
    return render_template('recommendations.html', recommended_movies=recommended_movies)

if __name__ == '__main__':
    app.run(debug=True)
