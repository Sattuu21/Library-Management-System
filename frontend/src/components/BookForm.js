import React, { useState, useEffect } from 'react';

const BookForm = ({ addBook, updateBook, currentBook }) => {
  const [book, setBook] = useState({ id: null, name: '', author: '', type: '' });

  useEffect(() => {
    if (currentBook) {
      setBook(currentBook);
    } else {
      setBook({ id: null, name: '', author: '', type: '' });
    }
  }, [currentBook]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setBook({ ...book, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (book.id) {
      console.log(`Updating book with ID: ${book.id}`);
      updateBook(book.id, book);
    } else {
      console.log(`Creating new book`);
      addBook(book);
    }
    setBook({ id: null, name: '', author: '', type: '' });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="form-group">
        <input
          type="text"
          name="name"
          value={book.name}
          onChange={handleChange}
          className="form-control"
          placeholder="Book Name"
          required
        />
      </div>
      <div className="form-group">
        <input
          type="text"
          name="author"
          value={book.author}
          onChange={handleChange}
          className="form-control"
          placeholder="Author"
          required
        />
      </div>
      <div className="form-group">
        <input
          type="text"
          name="type"
          value={book.type}
          onChange={handleChange}
          className="form-control"
          placeholder="Type"
          required
        />
      </div>
      <button type="submit" className="btn btn-primary">
        {book.id ? 'Update Book' : 'Add Book'}
      </button>
    </form>
  );
};

export default BookForm;
