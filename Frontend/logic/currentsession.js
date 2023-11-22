class CurrentSession {
    constructor() {
      if (CurrentSession.instance) {
        return CurrentSession.instance;
      }
  
      // Initialize your session properties here
      this.isLoggedIn = false;
      this.user = null;
  
      // Mark this instance as the singleton instance
      CurrentSession.instance = this;
  
      return this;
    }
  
    displayUser() {
      if (this.isLoggedIn && this.user) {
        console.table(this.user)
      } else {
        console.log('Not logged in.');
      }
    }

    login(user){
        this.isLoggedIn = true;
        this.user = user;
    }
  }

const currentSession = new CurrentSession();

export default currentSession;