const time = 3000
context('Access in the all forms to reservation page',()=>{

    it('Click the sport button to redirect into the reservation page and search ',()=>{
        cy.visit('http://localhost:4200/#/')
        cy.get(':nth-child(1) > .btn-sport').click()
        cy.wait(time)
        cy.get(':nth-child(1) > .nav-link').click()
        cy.wait(time)
        cy.get('.navbar-nav > :nth-child(2) > .nav-link').click()
        cy.get('.input-group-text').click()
        cy.wait(time-1500)
        cy.get('.input-group > :nth-child(2)').select("21:00")
        cy.wait(time-1500)
        cy.get('.hfin').select("22:00")
        cy.wait(time-1500)
        cy.get('.btn-search-reservation').click()
        cy.wait(time-1000)
        cy.get(':nth-child(1) > #confirm').click()
        cy.wait(time-1000)
        cy.get('.btn-success').click()
    })

    it('Do the login',()=>{
        cy.get('#typeEmailX-2').type('ejemplo@ejemplo.com')
        cy.wait(time-1800)
        cy.get('#typePasswordX-2').type('ejemplo')
        cy.wait(time)
        cy.get(':nth-child(1) > section > .container > .row > .col-12 > .card > .card-body > .btn').click()
        cy.wait(time)
    })

    //Can't show the Stripe because Cypress not support external testing
    it('Reservation the court ',()=>{
        cy.get('.navbar-nav > :nth-child(2) > .nav-link').click()
        cy.get('.input-group-text').click()
        cy.wait(time-1500)
        cy.get('.input-group > :nth-child(2)').select("21:00")
        cy.wait(time-1500)
        cy.get('.hfin').select("22:00")
        cy.wait(time-1500)
        cy.get('.btn-search-reservation').click()
        cy.wait(time-1000)
        cy.get(':nth-child(1) > #confirm').click()
        cy.wait(time-1000)
        cy.get('.btn-success').click()
    })
})