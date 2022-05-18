const time = 3000
context('Login with all modes and show the profile and user information',()=>{

    //Can't show the Social mode because Cypress not support this
    it('Social login and profile information',()=>{
        cy.visit('http://localhost:4200/#/')
        cy.wait(time)
        cy.get(':nth-child(6) > .nav-link').click()
        cy.wait(time-1800)
        cy.get(':nth-child(1) > .firebaseui-idp-button')
        cy.wait(time)
        cy.get(':nth-child(1) > .nav-link')
        cy.wait(time-2000)
        cy.get(':nth-child(6) > .nav-link').click()
        cy.wait(time-1800)
    })

    it('Do the login',()=>{
        cy.get('#typeEmailX-2').type('ejemplo@ejemplo.com')
        cy.wait(time-1800)
        cy.get('#typePasswordX-2').type('ejemplo')
        cy.wait(time)
        cy.get(':nth-child(1) > section > .container > .row > .col-12 > .card > .card-body > .btn').click()
        cy.wait(time)
    })

    it('Show user information',()=>{
        cy.get('#navbarDropdown').click()
        cy.wait(time-1500)
        cy.get(':nth-child(1) > .dropdown-item').click()
        //It has not been possible to access the profile and dashboard, the route returns to home for no reason
        cy.get('.navbar-nav > :nth-child(3) > .nav-link').click()
        cy.wait(time+2000)
        cy.get('.navbar-nav > :nth-child(4) > .nav-link').click()
        cy.wait(time)
        cy.get('footer').scrollIntoView({ duration: time })
    })
})