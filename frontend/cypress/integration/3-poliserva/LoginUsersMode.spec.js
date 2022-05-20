const time = 3000
context('Login with all modes, show the profile and user information and change language',()=>{

    it('Accept cookies', () => {
        cy.visit('http://localhost:4200/#/')
        cy.get('.buttons__cta > button').click()
    })

    // Can't show the Social Login mode because Cypress does not support external windows
    it('Social login and profile information',() => {
        cy.wait(time)
        cy.get(':nth-child(7) > .nav-link').click()
        cy.get(':nth-child(1) > .firebaseui-idp-button')
        cy.get(':nth-child(1) > .nav-link')
        cy.get(':nth-child(6) > .nav-link').click()
        cy.wait(time-1800)
    })

    it('User login',() => {
        cy.get('#typeEmailX-2').type('ejemplo@ejemplo.com')
        cy.wait(time-1800)
        cy.get('#typePasswordX-2').type('ejemplo')
        cy.get(':nth-child(1) > section > .container > .row > .col-12 > .card > .card-body > .btn').click()
    })

    it('Show user information and access to dashboard',() => {
        cy.wait(time-800)
        cy.get('#navbarDropdown').click()
        cy.get(':nth-child(1) > .dropdown-item').click()
        // It has not been possible to access the profile and dashboard, it seems Vue Guards don't let Cypress access to these routes    
    })


    it('Change Language', () => {
        cy.get(':nth-child(1) > .nav-link').contains('Home')
        cy.get('.form-select').select('Español')
        cy.get(':nth-child(1) > .nav-link').contains('Inicio')
    })

})