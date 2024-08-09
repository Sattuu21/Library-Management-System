import React from 'react';

const BookList = ({ books, deleteBook, setCurrentBook, selectBook }) => {
  return (
    <div>
      <h2>Your books</h2>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>Name</th>
            <th>Author</th>
            <th>Type</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {books.map((book) => (
            <tr key={book.id}>
              <td>{book.name}</td>
              <td>{book.author}</td>
              <td>{book.type}</td>
              <td>
                <button className="btn btn-info btn-sm" onClick={() => selectBook(book)}>View</button>
                <button className="btn btn-warning btn-sm" onClick={() => setCurrentBook(book)}>Update</button>
                <button className="btn btn-danger btn-sm" onClick={() => {
                  console.log(`Deleting book with id: ${book.id}`);
                  deleteBook(book.id);
                }}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BookList;
