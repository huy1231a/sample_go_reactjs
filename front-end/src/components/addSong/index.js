import React, { useState } from 'react';
import axios from 'axios';

const AddSong = () => {
    const [name, setName] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        axios.post('/songs', { Name: name })
            .then(response => {
                console.log(response);
            })
            .catch(error => {
                console.log(error);
            });
    }

    return (
        <div>
            <h1>Add Song</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={name} onChange={e => setName(e.target.value)} />
                </label>
                <button type="submit">Add</button>
            </form>
        </div>
    );
}

export default AddSong;
