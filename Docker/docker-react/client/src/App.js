import React from 'react'

import Header from './Components/Header'

export default function App() {
    return (
		<div className=' h-screen bg-no-repeat bg-center' style={{ backgroundImage: 'url("images/docker-img.png")'}} >
			<Header/>
		</div>
    )
}
