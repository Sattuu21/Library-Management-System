import React, { useState, useEffect } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import BookList from './components/BookList';
import BookForm from './components/BookForm';
import BookDetail from './components/Bookdetail';

const App = () => {
  const [books, setBooks] = useState([]);
  const [currentBook, setCurrentBook] = useState(null);
  const [selectedBook, setSelectedBook] = useState(null);

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async () => {
    try {
      const response = await fetch('http://localhost:9010/book/');
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setBooks(data);
    } catch (error) {
      console.error('Error fetching books:', error);
    }
  };

  const addBook = async (book) => {
    try {
        const response = await fetch('http://localhost:9010/book/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(book)
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log('Book added:', data);

        // Update the state with the newly added book
        setBooks([...books, data]);
    } catch (error) {
        console.error('Error adding book:', error);
    }
};


  const updateBook = async (id, updatedBook) => {
    console.log(`Updating book with id: ${id}`); // Log the ID
    if (id === undefined || id === null) {
      console.error("Invalid book ID");
      return;
    }
    try {
      const response = await fetch(`http://localhost:9010/book/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(updatedBook)
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setBooks(books.map(book => (book.id === id ? data : book)));
    } catch (error) {
      console.error('Error updating book:', error);
    }
  };

  const deleteBook = async (id) => {
    console.log(`Deleting book with id: ${id}`);
    if (id === undefined || id === null) {
      console.error("Invalid book ID");
      return;
    }
    try {
      const response = await fetch(`http://localhost:9010/book/${id}`, {
        method: 'DELETE'
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      setBooks(books.filter(book => book.id !== id));
    } catch (error) {
      console.error('Error deleting book:', error);
    }
  };

  const fetchBookById = async (id) => {
    console.log(`Fetching book with id: ${id}`);
    if (id === undefined || id === null) {
      console.error("Invalid book ID");
      return;
    }
    try {
      const response = await fetch(`http://localhost:9010/book/${id}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setSelectedBook(data);
    } catch (error) {
      console.error('Error fetching book by id:', error);
    }
  };

  const selectBook = (book) => {
    console.log(`Selecting book with id: ${book.id}`);
    fetchBookById(book.id);
    setCurrentBook(book);
  };

  const clearSelectedBook = () => {
    setSelectedBook(null);
  };

  return (
    <div className="container">
      <h1 className="my-4">Book Library</h1>
      {selectedBook ? (
        <BookDetail book={selectedBook} clearSelectedBook={clearSelectedBook} />
      ) : (
        <>
          <BookForm addBook={addBook} updateBook={updateBook} currentBook={currentBook} />
          <BookList books={books} deleteBook={deleteBook} setCurrentBook={setCurrentBook} selectBook={selectBook} />
        </>
      )}
    </div>
  );
};

export default App;
