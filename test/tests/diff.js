const comparePdf = require("compare-pdf");
const chai = require("chai");
const expect = chai.expect;
const config = require("../config/config");

describe("Compare different Pdf Tests in Mocha + Chai", () => {
    it("Compare different PDFs by image", async () => {
        const ComparePdf = new comparePdf(config);
        let comparisonResults = await ComparePdf.actualPdfFile("new.pdf")
            .baselinePdfFile("baseline.pdf")
            .compare("byImage");
        expect(comparisonResults.status).to.equal("failed");
        expect(comparisonResults.message).to.equal("new.pdf is not the same as baseline.pdf compared by their images.");
        expect(comparisonResults.details).to.not.be.null;
    });

    it("Compare different PDFs by base64", async () => {
        const ComparePdf = new comparePdf(config);
        let comparisonResults = await ComparePdf.actualPdfFile("new.pdf")
            .baselinePdfFile("baseline.pdf")
            .compare("byBase64");
        expect(comparisonResults.status).to.equal("failed");
        expect(comparisonResults.message).to.equal("new.pdf is not the same as baseline.pdf compared by their base64 values.");
        expect(comparisonResults.details).to.not.be.null;
    });
});