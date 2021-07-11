import React from 'react'


export default class basic_class_comp extends React.Component{
  constructor(props) {
		super(props)
    this.state = {}
  }

  componentDidCatch() { }
	componentDidMount = () => {}

  render(){
    return <h1> rename this comp </h1>
  }
}