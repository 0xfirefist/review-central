import AppBar from './components/appbar'
import ReviewBoard from './components/reviewboard'

function Page() {
    return (
      <div>
        <br />
        <h1>WELCOME TO REVIEW CENTRAL</h1>
        <br />
        <AppBar />
        <h2>THE REVIEW BOARD</h2>
        <br />
        <br />
        <ReviewBoard />
      </div>
    )
}

export default Page