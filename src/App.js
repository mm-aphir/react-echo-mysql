import React, { Component } from 'react';
import './App.css';
import Axios from 'axios'

class App extends Component {

  constructor(props) {
    super(props)

    this.state = {  member: [] , gender: [] , 
                    firstname: '', 
                    lastname: '', 
                    gender_id: '',
                    member_id: '', 
                    age: '' 
                  }
    this.handleInput = this.handleInput.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }
  componentDidMount() {
    { this.getMember() }
    { this.getGender() }
  }

  getMember = () => {
    var dataMember = []
    Axios.get('http://localhost:1323/member').then(result => {
        console.log(result.data)
        result.data.forEach(item => {
          dataMember.push(item)
        })
        this.setState({ member: dataMember })
    })
  }

  getGender = () => {
    var dataGender = []
    Axios.get('http://localhost:1323/gender').then(result => {
        console.log(result.data)
        result.data.forEach(item => {
          dataGender.push(item)
        })
        this.setState({ gender: dataGender })
    })
  }

  handleInput(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;
        this.setState({
            [name]: value,
        });
  }

  handleSubmit = event => {
    event.preventDefault();
    Axios.post(`http://localhost:1323/member`, {
        firstname: this.state.firstname,
        lastname: this.state.lastname,
        age: parseInt(this.state.age),
        gender_id: parseInt(this.state.gender_id)
    })
        .then(res => {
            console.log(res);
            console.log(res.data);
            console.log('Member Added!!');
        })
    console.dir(this.state);
    window.location.reload();
  }

  handleDelete(member_id) {
    Axios.delete(`http://localhost:1323/member/${member_id}`)
        .then(() => {
          console.log('Member Deleted!!');
        });
    window.location.reload();
  }

  render () {
    return (
      <div>
        <form className="toolbar"> Member </form>
        <center>
          <form onSubmit={this.handleSubmit}>
            <br/>
            <input type="text" name="firstname" placeholder = "Firstname" value={this.state.firstname} onChange={this.handleInput}/>
            <br/>
            <input type="text" name="lastname" placeholder = "Lastname" value={this.state.lastname} onChange={this.handleInput}/>
            <br/>
            <input type="text" name="age" placeholder = "Age" value={this.state.age} onChange={this.handleInput}/>
            <br/>
                
            <select name="gender_id" value={this.state.gender_id} onChange={this.handleInput}>
              <option disabled={true} value="">Select Gender</option>
                { this.state.gender.map((item, index) => (
                  <option value={item.gender_id} key={index} >{item.genderName}</option>
                )) }
            </select>
            
            <br/>
            <button type="submit">Submit</button>
          </form>
          <hr/>

          <table border='1' width='80%' cellSpacing="0">
            <tbody>
              <tr>
                <th>MemberID</th>
                <th>Firstname</th>
                <th>Lastname</th>
                <th>Age</th>
                <th>Gender</th>
                <th>Delete</th>
              </tr>
              {this.state.member.map((item, index) => (
                <tr key={index}>          
                  <td>{item.member_id}</td>
                  <td>{item.firstName}</td>
                  <td>{item.lastName}</td>
                  <td>{item.age}</td>
                  <td>{item.Gender.genderName}</td>
                  <td><button onClick={()=>{this.handleDelete(item.member_id)}}>Delete</button></td>
                </tr>
              ))}
            </tbody>
          </table>
        </center>
      </div>
    );
  }
}

export default App;
