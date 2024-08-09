import React from 'react';

const BookDetail = ({ book, clearSelectedBook }) => {
  return (
    <div>
      <h2>Book Detail</h2>
      <div className="card">
        <div className="card-body">
          <h5 className="card-title">{book.name}</h5>
          <h6 className="card-subtitle mb-2 text-muted">{book.author}</h6>
          <p className="card-text">{book.type}</p>
          <button className="btn btn-secondary" onClick={clearSelectedBook}>Back</button>
        </div>
      </div>
    </div>
  );
};

export default BookDetail;
