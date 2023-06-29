import React, { useState, useEffect } from 'react';
// import axios from 'axios';

const SongList = () => {
    const [products, setProducts] = useState([]);
    const [songs, setSong] = useState([]);

    // useEffect(() => {
    //     axios.get('http://localhost:8080/songs')
    //         .then(response => {
    //             // console.log("response" , response)
    //             // setSongs(response.data);
    //             const list = response.data
    //             const songlist = list.map(song => {
    //                 return {
    //                     ...song
    //                 }
    //             })
    //             setSongs(songlist)
    //         })
    //         .catch(error => {
    //             console.log(error);
    //         });
    // }, []);

    useEffect(() => {
        fetch('http://localhost:8000/api/products')
          .then(response => response.json())
          .then(data => setProducts(data));
      }, []);

    useEffect (() => {
        fetch('http://localhost:8000/api/songs')
        .then(response => response.json())
        .then(data => setSong(data))
    })

    return (
        <div>
            <h1>Product List</h1>
            {products.map(product => (
            <ul>
                <li className='li'>
                    Name: {product.name}
                </li>
                <li>
                    Price: {product.price}
                </li>
            </ul>
            ))}

            <h1>Song List</h1>
            {songs.map(song => (
            <ul>
                <li className='li'>
                    Name: {song.name}
                </li>
                <li>
                    Single Name: {song.single}
                </li>
            </ul>
            ))}
      </div>
    );
}

export default SongList;
